package book

import (
	"encoding/json"
	"io"
	"strconv"

	"github.com/gin-gonic/gin"
	. "iknowbook.com/data"
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
	insertProducts := func(bs []Book) error {
		res := ctr.productRepo.Insert(getProductsFromBooks(bs))
		_, err := res.RowsAffected()
		return err
	}
	insertBooks := func(bs []Book) {
		ctr.bookRepo.Insert(bs)
		ctx.JSON(200, bs)
	}
	insertFromBody := func(body []byte) {
		var bs []Book
		err := json.Unmarshal(body, &bs)
		if err == nil {
			err := insertProducts(bs)
			if err != nil {
				panic(err)
			}
			insertBooks(bs)
		} else {
			ctx.JSON(400, map[string]string{
				"Response": "The body should be list of books in json format.",
				"Error":    err.Error(),
			})
		}
	}

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		insertFromBody(body)
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
	updateFromBody := func(body []byte) {
		var books []Book
		err := json.Unmarshal(body, &books)
		if err == nil {
			ctr.bookRepo.Update(books)
			ctx.JSON(200, map[string]string{
				"Response": "Update successful",
			})
		} else {
			ctx.JSON(400, map[string]string{
				"Response": "The body should be list of products in json format.",
			})
		}
	}

	body, err := io.ReadAll(ctx.Request.Body)
	if err == nil {
		updateFromBody(body)
	} else {
		panic(err)
	}
}
