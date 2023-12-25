package user

import (
	"database/sql"
	"encoding/json"
	"io"
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
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Response": "Not a user Error:" + err.Error(),
				"Body":     string(body),
			})
		}
	}

	body, err := io.ReadAll(ctx.Request.Body)
	if err == nil {
		handleBody(body, operation, ctx)
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Response": "Reading request body failed. ERROR: " + err.Error(),
		})
	}
}

func (ser UserService) InsertUser(ctx *gin.Context, user User) {
	checkUserExist := func(result sql.Result, us User) {
		affectedRows, err := result.RowsAffected()
		if err == nil {
			if affectedRows > 0 {
				ctx.JSON(http.StatusOK, us)
			} else {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"Response": "User exist",
				})
			}
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"Response": "ERROR: " + err.Error(),
			})
		}
	}

	hashUserPasswordAndInsert := func(us User) {
		bytes, err := bcrypt.GenerateFromPassword([]byte(us.Password), 0)
		us.Password = string(bytes)

		if err == nil {
			result := ser.repo.Insert(us)
			checkUserExist(result, us)
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"Response": "Fail to hash password. ERROR: " + err.Error(),
			})
		}
	}

	hashUserPasswordAndInsert(user)
}

func (ser UserService) Insert(ctx *gin.Context) {
	insertUser := func(user User) {
		ser.InsertUser(ctx, user)
	}
	readAndHandleRequestBody(ctx, insertUser)
}

func (ser UserService) FindUserInfo(ctx *gin.Context) {
	handleUser := func(us User) {
		comparePassword := func(user User, pw string, ctx *gin.Context) {
			err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pw))
			if err == nil {
				ctx.JSON(http.StatusOK, user)
			} else {
				ctx.JSON(http.StatusUnauthorized, gin.H{
					"Response": "Password incorrect",
				})
			}
		}

		users := ser.repo.FindUserInfo(us)

		if len(users) > 0 {
			comparePassword(users[0], us.Password, ctx)
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Response": "user not found",
			})
		}
	}

	readAndHandleRequestBody(ctx, handleUser)
}

func (serv UserService) FindUserId(ctx *gin.Context) {
	readAndHandleRequestBody(ctx, func(user User) {
		users := serv.repo.FindUserInfo(user)
		if len(users) > 0 {
			ctx.JSON(http.StatusOK, gin.H{
				"id": users[0].Id,
			})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Response": "user not found",
			})
		}
	})
}

func (ser UserService) UpdateUserAuth(ctx *gin.Context) {
	handleUser := func(usr User) {
		result := ser.repo.UpdateUserAuth(usr)
		affectedRows, err := result.RowsAffected()
		if err == nil {
			if affectedRows > 0 {
				ctx.JSON(http.StatusOK, gin.H{
					"Response": "User Authorized",
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"Response": "User Authorization failed",
				})
			}
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"Response": "ERROR: " + err.Error(),
			})
		}
	}
	readAndHandleRequestBody(ctx, handleUser)
}
