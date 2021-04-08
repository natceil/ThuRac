package routes

import (
	"errors"
	"github.com/alexedwards/scs"
	"github.com/casbin/casbin"
	"golangSql/m/models"
	"log"
	"net/http"
)

var session *scs.Session

// Authorizer is a middleware for authorization
func Authorizer(e *casbin.Enforcer, users models.Users) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			role,err := session.GetString("role")
			if err != nil {
				WriteError(http.StatusInternalServerError, "ERROR", w, err)
				return
			}
			if role == "" {
				role = "anonymous"
			}
			// if it's a member, check if the user still exists
			if role == "member" {
				uid,err := session.GetInt("userId")
				if err != nil {
					WriteError(http.StatusInternalServerError, "ERROR", w, err)
					return
				}
				exists := users.Exists(uid)
				if !exists {
					WriteError(http.StatusForbidden, "FORBIDDEN", w, errors.New("user does not exist"))
					return
				}
			}
			// if it's a admin, check if the user still exists
			if role == "admin" {
				uid, err := session.GetInt("userId")
				if err != nil {
					WriteError(http.StatusInternalServerError, "ERROR", w,err)
					return
				}
				exists := users.Exists(uid)
				if !exists {
					WriteError(http.StatusForbidden, "FORBIDDEN", w, errors.New("user does not exist"))
					return
				}
			}
			// casbin enforce
			res, err := e.EnforceSafe(role, r.URL.Path, r.Method)
			if err != nil {
				WriteError(http.StatusInternalServerError, "ERROR", w, err)
				return
			}
			if res {
				next.ServeHTTP(w, r)
			} else {
				WriteError(http.StatusForbidden, "FORBIDDEN", w, errors.New("unauthorized"))
				return
			}
		}
		return http.HandlerFunc(fn)
	}
}

func WriteError(status int, message string, w http.ResponseWriter, err error) {
	log.Print("ERROR: ", err.Error())
	w.WriteHeader(status)
	w.Write([]byte(message))
}