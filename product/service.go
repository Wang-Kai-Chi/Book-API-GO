package product

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	. "iknowbook.com/app/data"
)

type ProductService struct {
	repo ProductRepository
}

func NewProductService(repo ProductRepository) ProductService {
	return ProductService{repo: repo}
}

func handleProductsFromContext(operation func([]Product) sql.Result, ctx *gin.Context) {
	handleBody := func(body []byte, operation func([]Product) sql.Result, ctx *gin.Context) {
		var ps []Product
		err := json.Unmarshal(body, &ps)
		if err == nil {
			operation(ps)
			ctx.JSON(http.StatusOK, ps)
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Response": "not products",
			})
		}
	}

	body, err := io.ReadAll(ctx.Request.Body)
	if err == nil {
		handleBody(body, operation, ctx)
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Response": "Reading request body failed. ERROR:" + err.Error(),
		})
	}
}

func (ctr ProductService) Insert(ctx *gin.Context) {
	handleProductsFromContext(ctr.repo.Insert, ctx)
}

func getPriceRange(ctx *gin.Context) (int, int) {
	min, err := strconv.Atoi(ctx.Query("min"))
	if err != nil {
		min = 0
	}
	max, err := strconv.Atoi(ctx.Query("max"))
	if err != nil {
		max = 0
	}
	return min, max
}

func (ctr ProductService) QueryWithPriceRange(ctx *gin.Context) {
	min, max := getPriceRange(ctx)
	ctx.JSON(http.StatusOK, ctr.repo.QueryWithPriceRange(min, max))
}

func (ctr ProductService) QueryByBarcode(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ctr.repo.QueryByBarcode(ctx.Param("barcode")))
}

func (ctr ProductService) Update(ctx *gin.Context) {
	handleProductsFromContext(ctr.repo.Update, ctx)
}

func (ctr ProductService) Delete(ctx *gin.Context) {
	handleProductsFromContext(ctr.repo.Delete, ctx)
}

func (ctr ProductService) QueryByConditions(ctx *gin.Context) {
	min, max := getPriceRange(ctx)
	product := Product{
		Product_title: "%" + ctx.DefaultQuery("title", "%") + "%",
		Publisher:     "%" + ctx.DefaultQuery("publisher", "%") + "%",
	}
	ctx.JSON(http.StatusOK, ctr.repo.QueryByConditions(min, max, product))
}

func (serv ProductService) MaxPrice(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"MaxProductPrice": serv.repo.MaxPrice()})
}

func (serv ProductService) QueryNewest(ctx *gin.Context) {
	ran, err := strconv.Atoi(ctx.Param("range"))
	if err == nil {
		ctx.JSON(http.StatusOK, serv.repo.QueryNewest(ran))
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Response": "uninspected input",
		})
	}
}
