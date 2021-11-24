package jwt

import (
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const expire = time.Hour * 336

func getToken(req *http.Request) string {
	cleared := strings.Replace(req.Header.Get("Authorization"), " ", "", -1)
	return strings.Replace(cleared, "Bearer", "", -1)
}

func Create(secretKey []byte, id string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["jti"] = id
	claims["exp"] = time.Now().Add(expire).Unix()

	return token.SignedString(secretKey)
}

func Parse(secretKey []byte, req *http.Request) (*jwt.Token, error) {
	return jwt.Parse(getToken(req), func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
}
