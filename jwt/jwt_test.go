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
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MDEyMjk4MTQsInVzZXIiOiJ0ZXN0dXNlciJ9.rs9rV7z1so6fbR8gNFibc04-duwbwpStPzpAP_lT2GM"
	res := VerifyJWTToken([]byte(key), token)
	if !res {
		t.Fail()
	}
}
