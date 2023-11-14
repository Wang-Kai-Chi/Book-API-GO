package user

import (
	"testing"

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
		users := []User{
			{
				Name:     "testuser2",
				Email:    "test2@mail.com",
				Phone:    "12345",
				Password: "testPassword2",
			},
		}
		rs := u.Insert(users)
		t.Log(rs)
	} else {
		t.Fatal(err)
	}
}
