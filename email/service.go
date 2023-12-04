package email

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EmailService struct{}

func NewEmailService() EmailService {
	return EmailService{}
}

type MailRecevier struct {
	Receiver string
}

func (serv EmailService) SendVerificationEmail(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		panic(err)
	}
	var mr MailRecevier
	err = json.Unmarshal(body, &mr)
	if err != nil {
		panic(err)
	}
	form := EMail{
		Sender:   "ericwangcatch@gmail.com",
		Subject:  "Iknowbook email Verification",
		HTMLBody: "Verify your Iknowbook acount email",
	}
	form.Receiver = mr.Receiver
	err = SendMail(form)
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
