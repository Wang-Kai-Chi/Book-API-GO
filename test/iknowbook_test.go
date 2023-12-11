package test

import (
	"testing"
)

func TestMustFindUserMenu(t *testing.T) {
	app := NewIknowBook()
	page := app.Start()
	open := app.MustFindUserMenu(page)
	if !open {
		t.Fatal("Can't find menu")
	}
}

func TestMustOpenMenu(t *testing.T) {
	app := NewIknowBook()
	page := app.Start()
	open := app.MustOpenUserMenu(page)
	if !open {
		t.Fatal("User menu didn't open")
	}
}

func TestLoginLogout(t *testing.T) {
	app := NewIknowBook()
	page := app.Start()
	app.Login(page)
	app.Logout(page)
}
