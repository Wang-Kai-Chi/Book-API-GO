package user

import (
	"embed"

	"golang.org/x/crypto/bcrypt"
)

//go:embed key.txt
var embedKey embed.FS

func mustGetKey() []byte {
	key, err := embedKey.ReadFile("key.txt")
	if err != nil {
		panic(err)
	}
	return key
}

type UserAuth struct{}

func (ua UserAuth) MustGetAuth() []byte {
	bytes, err := bcrypt.GenerateFromPassword(mustGetKey(), 0)

	if err != nil {
		panic(err)
	}
	return bytes
}
