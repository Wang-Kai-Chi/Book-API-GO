package main

import (
	"github.com/gin-gonic/gin"
	. "iknowbook.com/app/jwt"
)

type JwtController struct {
	service JwtService
	group   *gin.RouterGroup
}

func NewJwtController(service JwtService, router *gin.Engine) JwtController {
	return JwtController{
		service: service,
		group:   router.Group("api/v1/jwt"),
	}
}

func (ctr JwtController) Run() {
	ctr.GetJwtToken()
	ctr.VerifyJWTToken()
}

func (ctr JwtController) GetJwtToken() {
	ctr.group.POST("/token", ctr.service.GetJwtToken)
}

func (ctr JwtController) VerifyJWTToken() {
	ctr.group.POST("/verify", ctr.service.VerifyJWTToken)
}
