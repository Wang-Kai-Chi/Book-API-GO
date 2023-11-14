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
		users := User{
			Id:       "",
			Name:     "testuser",
			Email:    "test@mail.com",
			Phone:    "12345",
			Password: "testPassword",
		}
		rs := u.Insert(users)
		rowCount, err := rs.RowsAffected()
		if err == nil {
			t.Log("effected rows:", rowCount)
		}
	} else {
		t.Fatal(err)
	}
}
