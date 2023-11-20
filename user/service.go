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
	body, err := io.ReadAll(ctx.Request.Body)
	if err == nil {
		var ps User
		err := json.Unmarshal(body, &ps)
		if err == nil {
			bytes, err := bcrypt.GenerateFromPassword([]byte(ps.Password), 0)
			ps.Password = string(bytes)
			if err != nil {
				panic(err)
			}
			operation(ps)
			ctx.JSON(200, ps)
		} else {
			ctx.JSON(400, map[string]string{
				"Response": "not products",
			})
		}
	} else {
		log.Fatal("Reading request body failed. ", err)
	}
}

func (ser UserService) Insert(ctx *gin.Context) {
	handleUsersFromContext(ser.repo.Insert, ctx)
}
