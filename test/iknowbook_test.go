package test

import (
	"testing"
)

func TestLoginLogout(t *testing.T) {
	app := NewIknowBook()
	page := app.Start()
	app.Login(page)
	app.Logout(page)
}
