package jwt

import (
	"embed"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

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

//go:embed key.txt
var embedKey embed.FS

func mustGetKey() []byte {
	key, err := embedKey.ReadFile("key.txt")
	if err != nil {
		panic(err)
	}
	return key
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
			token, err := GetJWTToken(mustGetKey(), user.Name)
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
	verifyToken := func(raw string) {
		token := strings.ReplaceAll(raw, "Bearer ", "")
		res := VerifyJWTToken(mustGetKey(), token)
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
	token := ctx.Request.Header["Authorization"]
	verifyToken(token[0])
}
