package jwt

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

const hmacSecret = "aaa"

func GetJWTToken(key []byte) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["exp"] = time.Now().Add(time.Hour).Unix()
	claims["authorized"] = true
	claims["user"] = "username"

	tokenString, err := token.SignedString(key)

	return tokenString, err
}

func VerifyJWTToken(key []byte, token string) {
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Fatal("Unauthorized")
		}
		return key, nil
	})
	if err != nil {
		log.Fatal(err)
	}
	if t.Valid {
		return
	} else {
		log.Fatal("Invalid token")
	}
}
