package user

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

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

func readAndHandleRequestBody(ctx *gin.Context, operation func(User)) {
	handleBody := func(body []byte, operation func(User), ctx *gin.Context) {
		var us User
		err := json.Unmarshal(body, &us)
		if err == nil {
			operation(us)
		} else {
			ctx.JSON(http.StatusBadRequest, map[string]string{
				"Response": "Not a user",
			})
		}
	}

	body, err := io.ReadAll(ctx.Request.Body)
	if err == nil {
		handleBody(body, operation, ctx)
	} else {
		log.Fatal("Reading request body failed. ", err)
	}
}

func (ser UserService) Insert(ctx *gin.Context) {
	hashUserPasswordAndInsert := func(us User) {
		bytes, err := bcrypt.GenerateFromPassword([]byte(us.Password), 0)
		us.Password = string(bytes)

		if err != nil {
			panic(err)
		}
		ser.repo.Insert(us)

		ctx.JSON(http.StatusOK, us)
	}

	readAndHandleRequestBody(ctx, hashUserPasswordAndInsert)
}

func (ser UserService) FindUserInfo(ctx *gin.Context) {
	handleUser := func(us User) {
		comparePassword := func(user User, pw string, ctx *gin.Context) {
			err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pw))
			if err == nil {
				ctx.JSON(http.StatusOK, user)
			} else {
				ctx.JSON(http.StatusUnauthorized, map[string]string{
					"Response": "Password incorrect",
				})
			}
		}

		users := ser.repo.FindUserInfo(us)

		if len(users) > 0 {
			comparePassword(users[0], us.Password, ctx)
		} else {
			ctx.JSON(http.StatusBadRequest, map[string]string{
				"Response": "user not found",
			})
		}
	}

	readAndHandleRequestBody(ctx, handleUser)
}
