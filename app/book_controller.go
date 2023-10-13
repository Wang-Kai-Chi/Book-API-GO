package main

import (
	"github.com/gin-gonic/gin"
	. "iknowbook.com/book"
)

type BookController struct {
	service BookService
	prefix  string
	router  *gin.Engine
}

func NewBookController(service BookService, router *gin.Engine) BookController {
	return BookController{
		service: service,
		prefix:  "/book",
		router:  router,
	}
}

func (ctr BookController) Run() {
	ctr.QueryWithLimit()
	ctr.QueryByCondition()

	ctr.Insert()
	ctr.Update()
}

func (ctr BookController) QueryWithLimit() {
	ctr.router.GET(ctr.prefix+"/query/:limit", ctr.service.QueryWithLimit)
}

func (ctr BookController) Insert() {
	ctr.router.POST(ctr.prefix+"/insert", ctr.service.Insert)
}

func (ctr BookController) QueryByCondition() {
	ctr.router.GET(ctr.prefix+"/query", ctr.service.QueryByConditions)
}

func (ctr BookController) Update() {
	ctr.router.PUT(ctr.prefix+"/update", ctr.service.Update)
}
