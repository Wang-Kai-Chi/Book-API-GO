package controller

import (
	"github.com/gin-gonic/gin"
	. "iknowbook.com/app/book"
	jwt "iknowbook.com/app/jwt"
)

type BookController struct {
	service BookService
	group   *gin.RouterGroup
}

func NewBookController(service BookService, router *gin.Engine) BookController {
	return BookController{
		service: service,
		group:   router.Group("api/v1/book"),
	}
}

func (ctr BookController) Run() {
	ctr.QueryByCondition()

	ctr.Insert()
	ctr.Update()
}

func (ctr BookController) Insert() {
	ctr.group.POST(
		"/insert",
		func(ctx *gin.Context) {
			jwt.VerifyBearerToken(ctx, ctr.service.Insert)
		},
	)
}

func (ctr BookController) QueryByCondition() {
	ctr.group.GET(
		"/query",
		func(ctx *gin.Context) {
			jwt.VerifyBearerToken(ctx, ctr.service.QueryByConditions)
		},
	)
}

func (ctr BookController) Update() {
	ctr.group.PUT(
		"/update",
		func(ctx *gin.Context) {
			jwt.VerifyBearerToken(ctx, ctr.service.Update)
		},
	)
}
