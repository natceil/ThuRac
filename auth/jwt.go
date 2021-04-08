package auth

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"strings"
	"time"
)

//create jwt
func Create(account string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["account"] = account
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix() //Token hết hạn sau 2 giờ
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("anhhoangdeptrai")))
}

func Extract(r *http.Request) string {
	bearerToken := r.Header.Get("Authorization")
	return strings.Split(bearerToken, " ")[1]
}

func ExtractUsernameFromToken(r *http.Request) (string, error) {
	var account string
	tokenString := Extract(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("anhhoangdeptrai")), nil
	})

	if err != nil {
		return account, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		account = fmt.Sprintf("%v", claims["account"])
	}

	return account, nil
}

func Verify(r *http.Request) error {
	tokenString := Extract(r)
	_, err := jwt.Parse(tokenString,
	func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("anhhoangdeptrai")), nil
	})
	if err != nil {
		return err
	}
	return nil
}

