package user

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
	. "iknowbook.com/app/data"
	. "iknowbook.com/app/db"
)

func TestQueryWithLimit(t *testing.T) {
	db, err := ConnectDB()

	if err == nil {
		u := NewUserRepository(db)
		users := u.QueryWithLimit(10)
		t.Log(users)
	} else {
		t.Fatal(err)
	}
}

func TestInsert(t *testing.T) {
	db, err := ConnectDB()

	if err == nil {
		u := NewUserRepository(db)
		users := User{
			Id:       "",
			Name:     "testuser2",
			Email:    "test2@mail.com",
			Phone:    "12345",
			Password: "testPassword",
		}

		bytes, err := bcrypt.GenerateFromPassword([]byte(users.Password), 14)
		users.Password = string(bytes)

		rs := u.Insert(users)
		rowCount, err := rs.RowsAffected()
		if err == nil {
			t.Log("effected rows:", rowCount)
		}
	} else {
		t.Fatal(err)
	}
}

func TestFindUserInfo(t *testing.T) {
	db, err := ConnectDB()

	if err == nil {
		u := NewUserRepository(db)
		user := User{
			Id:       "",
			Name:     "testuser2",
			Email:    "test2@mail.com",
			Phone:    "12345",
			Password: "testPassword",
		}
		t.Log(u.FindUserInfo(user))
	} else {
		t.Fatal(err)
	}
}
