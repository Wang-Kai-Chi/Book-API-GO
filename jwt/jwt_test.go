package jwt

import (
	"testing"

	. "iknowbook.com/app/data"
)

var key = "aaa"

func TestMustGetBearerToken(t *testing.T) {
	user := User{
		Name: "testuser",
	}

	token, err := GetJWTToken([]byte(key), user.Name)
	if err == nil {
		t.Log(token)
	} else {
		t.Log(err)
	}
}

func TestVerifyJWTToken(t *testing.T) {
	token := "noene"
	res := MustVerifyJWTToken([]byte(key), token)
	if res == false {
		t.Fatal()
	}
}
