package main

import (
	"github.com/gin-gonic/gin"
	. "iknowbook.com/app/email"
	jwt "iknowbook.com/app/jwt"
)

type EmailController struct {
	service EmailService
	group   *gin.RouterGroup
}

func NewEmailController(service EmailService, router *gin.Engine) EmailController {
	return EmailController{
		service: service,
		group:   router.Group("/api/v1/email"),
	}
}

func (ctr EmailController) Run() {
	ctr.SendVerificationEmail()
}

func (ctr EmailController) SendVerificationEmail() {
	ctr.group.POST(
		"/send",
		func(ctx *gin.Context) {
			jwt.VerifyBearerToken(ctx, ctr.service.SendVerificationEmail)
		},
	)
}
