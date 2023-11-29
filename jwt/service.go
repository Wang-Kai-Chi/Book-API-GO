package jwt

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	. "iknowbook.com/app/data"
	. "iknowbook.com/app/user"
)

type UnverifiedInfo struct {
	User
	Token string
}

type JwtService struct {
	userRepo UserRepository
}

func NewJwtService(userRepo UserRepository) JwtService {
	return JwtService{
		userRepo: userRepo,
	}
}

func handleBody(body []byte, operation func(UnverifiedInfo), ctx *gin.Context) {
	var us UnverifiedInfo
	err := json.Unmarshal(body, &us)
	if err == nil {
		operation(us)
	} else {
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"Response": "Not a user",
		})
	}
}

func readAndHandleRequestBody(ctx *gin.Context, operation func(UnverifiedInfo)) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err == nil {
		handleBody(body, operation, ctx)
	} else {
		log.Fatal("Reading request body failed. ", err)
	}
}

func (serv JwtService) GetJwtToken(ctx *gin.Context) {
	getToken := func(us UnverifiedInfo) {
		users := serv.userRepo.FindUserInfo(us.User)
		if len(users) > 0 {
			user := users[0]
			token, err := GetJWTToken([]byte(user.Email), user.Name)
			if err != nil {
				panic(err)
			}
			us.Token = token
			ctx.JSON(http.StatusOK, us)
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Response": "Unauthorized User",
			})
		}

	}
	readAndHandleRequestBody(ctx, getToken)
}

func (serv JwtService) VerifyJWTToken(ctx *gin.Context) {
	verifyToken := func(us UnverifiedInfo) {
		res := VerifyJWTToken([]byte(us.Email), us.Token)
		if res {
			ctx.JSON(http.StatusOK, gin.H{
				"Result": "Authorized",
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"Result": "Unauthorized",
			})
		}
	}
	readAndHandleRequestBody(ctx, verifyToken)
}