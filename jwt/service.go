package jwt

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	. "iknowbook.com/app/data"
	. "iknowbook.com/app/user"
)

type JwtService struct {
	userRepo UserRepository
}

func NewJwtService(userRepo UserRepository) JwtService {
	return JwtService{
		userRepo: userRepo,
	}
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

func (serv JwtService) GetJwtToken(ctx *gin.Context) {
	verifyUser := func(user User) {
		token, err := GetJWTToken([]byte(user.Auth), user.Name)
		if err == nil {
			ctx.JSON(http.StatusOK, gin.H{"Token": token})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"Response": "ERROR: " + err.Error(),
			})
		}
	}
	getToken := func(us User) {
		users := serv.userRepo.FindExactUserInfo(us)
		if len(users) > 0 {
			verifyUser(us)
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Response": "Unauthorized User",
			})
		}
	}
	readAndHandleRequestBody(ctx, getToken)
}

func VerifyBearerToken(ctx *gin.Context, authOp func(ctx *gin.Context)) {
	isVerified := func(bearer string, userAuth string) bool {
		token := strings.ReplaceAll(bearer, "Bearer ", "")
		res := MustVerifyJWTToken([]byte(userAuth), token)
		return res
	}

	handleVerification := func(key string) {
		bearers := ctx.Request.Header["Authorization"]

		if len(bearers) > 0 {
			bearer := bearers[0]
			if isVerified(bearer, key) {
				authOp(ctx)
			} else {
				if len(bearer) <= 0 {
					ctx.JSON(http.StatusBadRequest, gin.H{
						"Response": "No bearer token found in header.",
					})
				} else {
					ctx.JSON(http.StatusUnauthorized, gin.H{
						"Response": "Token expired or not jwt.",
					})
				}
			}
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Response": "No Authrization",
			})
		}
	}

	keys := ctx.Request.Header["Auth-Key"]
	if len(keys) > 0 {
		handleVerification(keys[0])
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Response": "No User auth key.",
		})
	}
}

func (serv JwtService) VerifyJWTToken(ctx *gin.Context) {
	VerifyBearerToken(ctx,
		func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"Response": "Authorized",
			})
		},
	)
}
