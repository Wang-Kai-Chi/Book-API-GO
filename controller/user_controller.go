package controller

import (
	"github.com/gin-gonic/gin"
	. "iknowbook.com/app/email"
	. "iknowbook.com/app/user"
)

type UserController struct {
	service UserService
	group   *gin.RouterGroup
}

func NewUserController(service UserService, router *gin.Engine) UserController {
	return UserController{
		service: service,
		group:   router.Group("api/v1/user"),
	}
}

func (ctr UserController) Run() {
	ctr.FindUserInfo()
	ctr.Insert()
	ctr.UpdateUserAuth()
}

func (ctr UserController) Insert() {
	ctr.group.POST("/insert",
		func(ctx *gin.Context) {
			VerifyUserEmail(ctx, ctr.service.InsertUser)
		},
	)
}

func (ctr UserController) FindUserInfo() {
	ctr.group.POST("/login", ctr.service.FindUserInfo)
}

func (ctr UserController) UpdateUserAuth() {
	ctr.group.POST("/auth", ctr.service.UpdateUserAuth)
}