package email

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	. "iknowbook.com/app/data"
)

type EmailService struct {
}

func NewEmailService() EmailService {
	return EmailService{}
}

func handleInternalError(err error, ctx *gin.Context) {
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Response": "ERROR: " + err.Error(),
		})
	}
}

func mustGetUserFromBody(ctx *gin.Context) User {
	body, err := io.ReadAll(ctx.Request.Body)
	handleInternalError(err, ctx)

	var us User
	err = json.Unmarshal(body, &us)
	handleInternalError(err, ctx)

	return us
}

func mustReadEmailForm() string {
	htmlF, err := os.ReadFile("../static/verify_email_form.html")
	if err != nil {
		panic(err)
	}
	return string(htmlF)
}

func (serv EmailService) SendVerificationEmail(ctx *gin.Context) {
	sendVerificationMail := func(receiver string) {
		form := EMail{
			Sender:   "ericwangcatch@gmail.com",
			Subject:  "Iknowbook email Verification",
			HTMLBody: string(mustReadEmailForm()),
		}
		form.Receiver = receiver

		err := SendMail(form)
		if err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"Response": "A verification mail has sended, please check your email and verify it.",
			})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Response": "Sending a verification mail failed, please try again.",
			})
		}
	}

	VerifyUserEmail(ctx,
		func(ctx *gin.Context, user User) {
			sendVerificationMail(user.Email)
		},
	)
}

func VerifyUserEmail(ctx *gin.Context, authOp func(*gin.Context, User)) {
	us := mustGetUserFromBody(ctx)

	err := VerifyEmail(us.Email)
	if err == nil {
		authOp(ctx, us)
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"Response": "Email verify failed. ERROR: " + err.Error(),
		})
	}
}

func (serv EmailService) VerifyEmail(ctx *gin.Context) {
	VerifyUserEmail(ctx,
		func(ctx *gin.Context, user User) {
			ctx.JSON(http.StatusOK, gin.H{
				"Response": "Email verified.",
				"Body":     user,
			})
		},
	)
}
