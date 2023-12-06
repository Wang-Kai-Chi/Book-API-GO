package main

import (
	"github.com/gin-gonic/gin"
	jwt "iknowbook.com/app/jwt"
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
	ctr.QueryByConditions()
	ctr.QueryByBarcode()
	ctr.QueryNewest()

	ctr.MaxPrice()

	ctr.Insert()
	ctr.Update()
	ctr.Delete()
}

func (ctr ProductController) Insert() {
	ctr.group.POST(
		"/insert",
		func(ctx *gin.Context) {
			jwt.VerifyBearerToken(ctx, ctr.service.Insert)
		},
	)
}

func (ctr ProductController) Update() {
	ctr.group.PUT(
		"/update",
		func(ctx *gin.Context) {
			jwt.VerifyBearerToken(ctx, ctr.service.Update)
		},
	)
}

func (ctr ProductController) Delete() {
	ctr.group.DELETE(
		"/delete",
		func(ctx *gin.Context) {
			jwt.VerifyBearerToken(ctx, ctr.service.Delete)
		},
	)
}

func (ctr ProductController) QueryByConditions() {
	ctr.group.GET(
		"/query/",
		func(ctx *gin.Context) {
			jwt.VerifyBearerToken(ctx, ctr.service.QueryByConditions)
		},
	)
}

func (ctr ProductController) QueryByBarcode() {
	ctr.group.GET(
		"/query/barcode/:barcode",
		func(ctx *gin.Context) {
			jwt.VerifyBearerToken(ctx, ctr.service.QueryByBarcode)
		},
	)
}

func (ctr ProductController) MaxPrice() {
	ctr.group.GET(
		"/maxprice",
		func(ctx *gin.Context) {
			jwt.VerifyBearerToken(ctx, ctr.service.MaxPrice)
		},
	)
}

func (ctr ProductController) QueryNewest() {
	ctr.group.GET(
		"/query/new/:range",
		ctr.service.QueryNewest,
	)
}
