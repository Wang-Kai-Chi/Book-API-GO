package book

import (
	"io"
	"strconv"

	"github.com/gin-gonic/gin"
	. "iknowbook.com/data"
	. "iknowbook.com/db"
	"iknowbook.com/product"
)

type BookService struct {
	bookRepo    BookRepository
	productRepo product.ProductRepository
}

func NewBookService(bookRepo BookRepository, productRepo product.ProductRepository) BookService {
	return BookService{
		bookRepo:    bookRepo,
		productRepo: productRepo,
	}
}

func (ctr BookService) QueryWithLimit(ctx *gin.Context) {
	limit, err := strconv.Atoi(ctx.Param("limit"))
	if err == nil {
		ctx.JSON(200, ctr.bookRepo.QueryWithLimit(limit))
	} else {
		ctx.JSON(400, map[string]string{
			"Response": "Please type number for limit.",
		})
	}
}

func (ctr BookService) Insert(ctx *gin.Context) {
	getProductsFromBooks := func(bs []Book) []Product {
		var temp []Product
		for _, v := range bs {
			temp = append(temp, v.Product)
		}
		return temp
	}
	body, err := io.ReadAll(ctx.Request.Body)
	bs, err := GetEntityFromBody[Book](body)
	if err == nil {
		res := ctr.productRepo.Insert(getProductsFromBooks(bs))
		_, err := res.RowsAffected()
		if err == nil {
			ctr.bookRepo.Insert(bs)
			ctx.JSON(200, bs)
		} else {
			ctx.JSON(400, map[string]string{
				"Response": "The body should be list of products in json format.",
			})
		}
	} else {
		panic(err)
	}
}

func (ctr BookService) QueryByConditions(ctx *gin.Context) {
	min, err := strconv.Atoi(ctx.Query("min"))
	if err != nil {
		min = 0
	}
	max, err := strconv.Atoi(ctx.Query("max"))
	if err != nil {
		max = 0
	}
	book := Book{
		Product: Product{
			Product_title: "%" + (ctx.DefaultQuery("title", "%")) + "%",
			Publisher:     "%" + (ctx.DefaultQuery("publisher", "%")) + "%",
		},
		Author:     "%" + (ctx.DefaultQuery("author", "%")) + "%",
		Translator: "%" + (ctx.DefaultQuery("translator", "%")) + "%",
		Language:   "%" + (ctx.DefaultQuery("language", "%")) + "%",
		Category:   "%" + (ctx.DefaultQuery("category", "%")) + "%",
	}

	ctx.JSON(200, ctr.bookRepo.QueryByConditions(min, max, book))
}

func (ctr BookService) Update(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)

	if err == nil {
		book, err := GetEntityFromBody[Book](body)
		if err == nil {
			ctr.bookRepo.Update(book)
			ctx.JSON(200, map[string]string{
				"Response": "Update successful",
			})
		} else {
			ctx.JSON(400, map[string]string{
				"Response": "The body should be list of products in json format.",
			})
		}
	} else {
		panic(err)
	}
}
