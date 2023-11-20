package main

import (
	"github.com/gin-gonic/gin"
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
	ctr.QueryWithLimit()

	ctr.FindUserInfo()
	ctr.Insert()
}

func (ctr UserController) QueryWithLimit() {
	ctr.group.GET("/query/:limit", ctr.service.QueryWithLimit)
}

func (ctr UserController) Insert() {
	ctr.group.POST("/insert", ctr.service.Insert)
}

func (ctr UserController) FindUserInfo() {
	ctr.group.POST("/login", ctr.service.FindUserInfo)
}
