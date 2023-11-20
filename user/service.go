package user

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	. "iknowbook.com/app/data"
)

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return UserService{repo: repo}
}

func (ser UserService) QueryWithLimit(ctx *gin.Context) {
	limit, err := strconv.Atoi(ctx.Param("limit"))
	if err == nil {
		ctx.JSON(200, ser.repo.QueryWithLimit(limit))
	} else {
		ctx.JSON(400, map[string]string{
			"Response": "unexpecting input",
		})
	}
}

func handleUsersFromContext(operation func(User) sql.Result, ctx *gin.Context) {
	handleBody := func(body []byte) {
		var us User
		err := json.Unmarshal(body, &us)
		if err == nil {
			operation(us)
			ctx.JSON(200, us)
		} else {
			ctx.JSON(400, map[string]string{
				"Response": "not user",
			})
		}
	}

	body, err := io.ReadAll(ctx.Request.Body)
	if err == nil {
		handleBody(body)
	} else {
		log.Fatal("Reading request body failed. ", err)
	}
}

func (ser UserService) Insert(ctx *gin.Context) {
	hashUserPasswordAndInsert := func(us User) sql.Result {
		bytes, err := bcrypt.GenerateFromPassword([]byte(us.Password), 0)
		us.Password = string(bytes)
		if err != nil {
			panic(err)
		}
		return ser.repo.Insert(us)
	}
	handleUsersFromContext(hashUserPasswordAndInsert, ctx)
}

func (ser UserService) FindUserInfo(ctx *gin.Context) {
	comparePassword := func(users []User, pw string, ctx *gin.Context) {
		err := bcrypt.CompareHashAndPassword([]byte(users[0].Password), []byte(pw))
		if err == nil {
			ctx.JSON(200, users)
		} else {
			ctx.JSON(401, map[string]string{
				"Response": "Password incorrect or User not exists.",
			})
		}
	}

	handleBody := func(body []byte) {
		var us User
		err := json.Unmarshal(body, &us)
		if err == nil {
			users := ser.repo.FindUserInfo(us)
			comparePassword(users, us.Password, ctx)
		} else {
			ctx.JSON(400, map[string]string{
				"Response": "not user",
			})
		}
	}

	body, err := io.ReadAll(ctx.Request.Body)
	if err == nil {
		handleBody(body)
	} else {
		log.Fatal("Reading request body failed. ", err)
	}
}
