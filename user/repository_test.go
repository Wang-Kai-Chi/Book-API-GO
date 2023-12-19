package user

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
	. "iknowbook.com/app/data"
	. "iknowbook.com/app/db"
)

func startDBOperateTest(operate func(UserRepository), t *testing.T) {
	db, err := ConnectDB()

	if err == nil {
		u := NewUserRepository(db)
		operate(u)
	} else {
		t.Fatal(err)
	}
}

func TestQueryWithLimit(t *testing.T) {
	startDBOperateTest(func(u UserRepository) {
		users := u.QueryWithLimit(10)
		t.Log(users)
	}, t)
}

func TestInsert(t *testing.T) {
	startDBOperateTest(func(u UserRepository) {
		users := User{
			Id:       "",
			Name:     "testuser3",
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
	}, t)
}

func TestFindUserInfo(t *testing.T) {
	startDBOperateTest(func(u UserRepository) {
		user := User{
			Id:       "",
			Name:     "testuserA",
			Email:    "testA@mail.com",
			Phone:    "a123456",
			Password: "testPassword",
		}
		t.Log(u.FindUserInfo(user))
	}, t)
}

func TestFindExactUserInfo(t *testing.T) {
	startDBOperateTest(func(u UserRepository) {
		user := User{
			Id:       "",
			Name:     "testuserA",
			Email:    "testA@mail.com",
			Phone:    "a123456",
			Password: "testpassword",
		}
		t.Log(u.FindExactUserInfo(user))
	}, t)
}

func TestUpdataUserAuth(t *testing.T) {
	startDBOperateTest(func(u UserRepository) {
		var auth UserAuth
		user := User{
			Id:   "27f7842f-475c-4ffd-b727-9892e0a0a6f0",
			Auth: string(auth.MustGetAuth()),
		}
		t.Log(u.UpdateUserAuth(user).RowsAffected())
	}, t)
}
