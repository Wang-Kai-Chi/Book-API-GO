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
			Name:     "testa",
			Email:    "testa@mail.com",
			Phone:    "",
			Password: "123",
			Auth:     "$2a$10$lvZxVD2sQn5mlcWpRCdTROmrGZIlyjvqxLp1euMX7pwm.Y0qxSNGm",
		}
		t.Log(u.FindExactUserInfo(user))
	}, t)
}

func TestUpdataUserAuth(t *testing.T) {
	startDBOperateTest(func(u UserRepository) {
		var auth UserAuth
		user := User{
			Id:   "d8f34f29-2f6b-4725-ba78-bbc871e6f5e6",
			Auth: string(auth.MustGetAuth()),
		}
		t.Log(u.UpdateUserAuth(user).RowsAffected())
	}, t)
}

func TestQueryById(t *testing.T) {
	startDBOperateTest(func(ur UserRepository) {
		id := "d8f34f29-2f6b-4725-ba78-bbc871e6f5e6"
		t.Log(ur.QueryById(id))
	}, t)
}
