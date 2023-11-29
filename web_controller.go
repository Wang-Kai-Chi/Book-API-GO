package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"iknowbook.com/app/book"
	. "iknowbook.com/app/db"
	"iknowbook.com/app/jwt"
	"iknowbook.com/app/product"
	"iknowbook.com/app/user"
)

type WebController struct{}

var (
	bookRepo    book.BookRepository
	productRepo product.ProductRepository
	userRepo    user.UserRepository
)

func mustInitRepos() {
	db, err := ConnectDB()
	if err == nil {
		bookRepo = book.NewBookRepository(db)
		productRepo = product.NewProductRepository(db)
		userRepo = user.NewUserRepository(db)
	} else {
		panic(err)
	}
}

func (w WebController) Init() {
	router := gin.New()
	router.Use(gin.Logger())
	router.SetTrustedProxies([]string{"127.0.0.1"})
	mustInitRepos()

	NewProductController(
		product.NewProductService(productRepo),
		router,
	).Run()

	NewBookController(
		book.NewBookService(bookRepo, productRepo),
		router,
	).Run()

	NewUserController(
		user.NewUserService(userRepo),
		router,
	).Run()

	NewJwtController(
		jwt.NewJwtService(userRepo),
		router,
	).Run()

	router.StaticFS("/static", http.Dir("static/"))
	index(router)

	addr := "localhost"
	port := ":8081"
	println("server start at " + addr + port)

	router.Run(addr + port)

}

func index(r *gin.Engine) {
	path := "./static/*.html"
	r.LoadHTMLGlob(path)
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
}
