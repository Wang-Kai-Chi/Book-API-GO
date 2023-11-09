package main

import (
	"github.com/gin-gonic/gin"
	. "iknowbook.com/app/book"
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
	ctr.QueryWithLimit()
	ctr.QueryByCondition()

	ctr.Insert()
	ctr.Update()
}

func (ctr BookController) QueryWithLimit() {
	ctr.group.GET("/query/:limit", ctr.service.QueryWithLimit)
}

func (ctr BookController) Insert() {
	ctr.group.POST("/insert", ctr.service.Insert)
}

func (ctr BookController) QueryByCondition() {
	ctr.group.GET("/query", ctr.service.QueryByConditions)
}

func (ctr BookController) Update() {
	ctr.group.PUT("/update", ctr.service.Update)
}
