package main

import (
	"github.com/gin-gonic/gin"
	. "iknowbook.com/product"
)

type ProductController struct {
	service ProductService
	prefix  string
	router  *gin.Engine
}

func NewProductController(service ProductService, router *gin.Engine) ProductController {
	return ProductController{
		service: service,
		prefix:  "/product",
		router:  router,
	}
}

func (ctr ProductController) Run() {
	ctr.QueryWithLimit()
	ctr.QueryByConditions()
	ctr.QueryByBarcode()

	ctr.Insert()
	ctr.Update()
	ctr.Delete()
}

func (ctr ProductController) QueryWithLimit() {
	ctr.router.GET(ctr.prefix+"/query/:limit", ctr.service.QueryWithLimit)
}

func (ctr ProductController) Insert() {
	ctr.router.POST(ctr.prefix+"/insert", ctr.service.Insert)
}

func (ctr ProductController) Update() {
	ctr.router.PUT(ctr.prefix+"/update", ctr.service.Update)
}

func (ctr ProductController) Delete() {
	ctr.router.PUT(ctr.prefix+"/delete", ctr.service.Delete)
}

func (ctr ProductController) QueryByConditions() {
	ctr.router.GET(ctr.prefix+"/query/", ctr.service.QueryByConditions)
}

func (ctr ProductController) QueryByBarcode() {
	ctr.router.GET(ctr.prefix+"/query/barcode/:barcode", ctr.service.QueryByBarcode)
}
