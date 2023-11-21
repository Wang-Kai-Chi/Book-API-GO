package jwt

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

func MustGetBearerToken() string {
	hmacSampleSecret := []byte("111")
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = jwt.MapClaims{
		"foo": "bar",
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	}
	tokenString, err := token.SignedString(hmacSampleSecret)

	if err != nil {
		log.Fatal(err)
	}

	return tokenString
}
