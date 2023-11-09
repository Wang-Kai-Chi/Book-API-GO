package main

import (
	"github.com/gin-gonic/gin"
	. "iknowbook.com/app/product"
)

type ProductController struct {
	service ProductService
	group   *gin.RouterGroup
}

func NewProductController(service ProductService, router *gin.Engine) ProductController {
	return ProductController{
		service: service,
		group:   router.Group("api/v1/product"),
	}
}

func (ctr ProductController) Run() {
	ctr.QueryWithLimit()
	ctr.QueryByConditions()
	ctr.QueryByBarcode()
	ctr.QueryNewest()

	ctr.MaxPrice()

	ctr.Insert()
	ctr.Update()
	ctr.Delete()
}

func (ctr ProductController) QueryWithLimit() {
	ctr.group.GET("/query/:limit", ctr.service.QueryWithLimit)
}

func (ctr ProductController) Insert() {
	ctr.group.POST("/insert", ctr.service.Insert)
}

func (ctr ProductController) Update() {
	ctr.group.PUT("/update", ctr.service.Update)
}

func (ctr ProductController) Delete() {
	ctr.group.DELETE("/delete", ctr.service.Delete)
}

func (ctr ProductController) QueryByConditions() {
	ctr.group.GET("/query/", ctr.service.QueryByConditions)
}

func (ctr ProductController) QueryByBarcode() {
	ctr.group.GET("/query/barcode/:barcode", ctr.service.QueryByBarcode)
}

func (ctr ProductController) MaxPrice() {
	ctr.group.GET("/maxprice", ctr.service.MaxPrice)
}

func (ctr ProductController) QueryNewest() {
	ctr.group.GET("/query/new/:range", ctr.service.QueryNewest)
}
