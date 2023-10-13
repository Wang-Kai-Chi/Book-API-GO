package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"iknowbook.com/book"
	. "iknowbook.com/db"
	"iknowbook.com/product"
)

type WebController struct{}

var (
	bookRepo    book.BookRepository
	productRepo product.ProductRepository
)

func mustInitRepos() {
	db, err := ConnectDB()
	if err == nil {
		bookRepo = book.NewBookRepository(db)
		productRepo = product.NewProductRepository(db)
	} else {
		panic(err)
	}
}

func (w WebController) Init() {
	router := gin.Default()
	router.SetTrustedProxies([]string{"http", "https"})
	mustInitRepos()

	NewProductController(
		product.NewProductService(productRepo),
		router,
	).Run()

	NewBookController(
		book.NewBookService(bookRepo, productRepo),
		router,
	).Run()

	index(router)

	addr := "localhost"
	port := ":8080"
	println("server start at " + addr + port)

	router.Run(addr + port)

}

func index(r *gin.Engine) {
	path := "./app/static/*.html"
	r.LoadHTMLGlob(path)
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
}
