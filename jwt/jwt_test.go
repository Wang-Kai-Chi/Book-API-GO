package jwt

import "testing"

func TestMustGetBearerToken(t *testing.T) {
	token := MustGetBearerToken()
	t.Log(token)
}
