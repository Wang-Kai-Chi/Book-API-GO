package email

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	. "iknowbook.com/app/data"
	. "iknowbook.com/app/user"
)

type EmailService struct {
	userRepo UserRepository
}

func NewEmailService(userRepo UserRepository) EmailService {
	return EmailService{
		userRepo: userRepo,
	}
}

func mustGetUserFromBody(ctx *gin.Context) User {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println(err)
	}

	var us User
	err = json.Unmarshal(body, &us)
	if err != nil {
		log.Println(err)
	}
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

	us := mustGetUserFromBody(ctx)

	users := serv.userRepo.FindUserInfo(us)

	if len(users) > 0 {
		user := users[0]
		sendVerificationMail(user.Email)
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Response": "User unregister.",
		})
	}
}

func (serv EmailService) VerifyEmail(ctx *gin.Context) {
	us := mustGetUserFromBody(ctx)

	err := VerifyEmail(us.Email)
	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"Response": "Email verified.",
		})
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"Response": "Email verify failed. ERROR: " + err.Error(),
		})
	}
}
