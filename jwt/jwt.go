package jwt

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

func GetJWTToken(key []byte, username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["exp"] = time.Now().Add(time.Hour).Unix()
	claims["authorized"] = true
	claims["user"] = username

	tokenString, err := token.SignedString(key)

	return tokenString, err
}

func MustVerifyJWTToken(key []byte, token string) bool {
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, jwt.ErrInvalidKeyType
		}
		return key, nil
	})
	if err != nil {
		log.Println(err)
		return false
	}
	if t.Valid {
		return true
	} else {
		log.Println("Invalid token")
		return false
	}
}
