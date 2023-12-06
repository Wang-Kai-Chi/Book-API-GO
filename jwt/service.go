package jwt

import (
	"embed"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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

func readAndHandleRequestBody(ctx *gin.Context, operation func(UnverifiedInfo)) {
	handleBody := func(body []byte, operation func(UnverifiedInfo), ctx *gin.Context) {
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
	body, err := io.ReadAll(ctx.Request.Body)
	if err == nil {
		handleBody(body, operation, ctx)
	} else {
		log.Fatal("Reading request body failed. ", err)
	}
}

func (serv JwtService) GetJwtToken(ctx *gin.Context) {
	verifyUser := func(user User, us UnverifiedInfo) {
		printVerifiedInfo := func() {
			token, err := GetJWTToken(mustGetKey(), user.Name)
			if err != nil {
				panic(err)
			}
			us.Token = token
			ctx.JSON(http.StatusOK, gin.H{"Bearer Token": us.Token})
		}

		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(us.Password))

		if err == nil {
			printVerifiedInfo()
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Response": "Unauthorized User",
			})
		}
	}
	getToken := func(us UnverifiedInfo) {
		users := serv.userRepo.FindUserInfo(us.User)
		if len(users) > 0 {
			verifyUser(users[0], us)
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Response": "Unauthorized User",
			})
		}

	}
	readAndHandleRequestBody(ctx, getToken)
}

func VerifyBearerToken(ctx *gin.Context, authOp func(ctx *gin.Context)) {
	isVerified := func(bearer string) bool {
		token := strings.ReplaceAll(bearer, "Bearer ", "")
		res := VerifyJWTToken(mustGetKey(), token)
		return res
	}
	bearers := ctx.Request.Header["Authorization"]
	if len(bearers) > 0 {
		if isVerified(bearers[0]) {
			authOp(ctx)
		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"Result": "Unauthorized",
			})
		}
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Response": "No bearer token found in header.",
		})
	}
}

func (serv JwtService) VerifyJWTToken(ctx *gin.Context) {
	VerifyBearerToken(ctx,
		func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"Result": "Authorized",
			})
		},
	)
}
