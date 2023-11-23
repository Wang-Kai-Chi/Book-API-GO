package jwt

import "testing"

var key = "aaa"

func TestMustGetBearerToken(t *testing.T) {
	token, err := GetJWTToken([]byte(key))
	if err == nil {
		t.Log(token)
	} else {
		t.Log(err)
	}
}

func TestVerifyJWTToken(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MDA3MzA0NDcsInVzZXIiOiJ1c2VybmFtZSJ9.hVgL842CnAMfXs7T_3w411P5f8gqPRF5-hX2QTkvJQY"
	VerifyJWTToken([]byte(key), token)
}
