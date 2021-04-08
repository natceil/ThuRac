package routes

import (
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"golangSql/m/auth"
	"golangSql/m/models"
	"html"
	"net/http"
	"strings"
)

//check user == user
func CheckJwt(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		err := auth.Verify(r)

		if err != nil {
			authen := &models.Authorized{
				Msg:    "Bạn chưa đăng nhập hoặc đã hết phiên sử dụng",
				Author: false,
			}
			userBytes, _ := json.Marshal(authen)

			w.Write(userBytes)
			return
		}
		next.ServeHTTP(w, r)
	})
}

//func prevent injection
func Santize(data string) string{
	data = html.EscapeString(strings.TrimSpace(data))
	return data
}

//func hash bcrypt pass
func Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}


