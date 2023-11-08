package product

import (
	"encoding/json"
	"io"
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

func (ctr ProductService) QueryWithLimit(ctx *gin.Context) {
	limit, err := strconv.Atoi(ctx.Param("limit"))
	if err == nil {
		ctx.JSON(200, ctr.repo.QueryWithLimit(limit))
	} else {
		ctx.JSON(400, map[string]string{
			"Response": "Please type number for limit.",
		})
	}
}

func (ctr ProductService) Insert(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err == nil {
		var ps []Product
		err := json.Unmarshal(body, &ps)
		if err == nil {
			ctr.repo.Insert(ps)
			ctx.JSON(200, ps)
		} else {
			ctx.JSON(400, map[string]string{
				"Response": "The body should be list of products in json format.",
			})
		}
	} else {
		panic(err)
	}
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
	ctx.JSON(200, ctr.repo.QueryWithPriceRange(min, max))
}

func (ctr ProductService) QueryByBarcode(ctx *gin.Context) {
	ctx.JSON(200, ctr.repo.QueryByBarcode(ctx.Param("barcode")))
}

func (ctr ProductService) Update(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err == nil {
		var ps []Product
		err := json.Unmarshal(body, &ps)
		if err == nil {
			ctr.repo.Update(ps)
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

func (ctr ProductService) Delete(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err == nil {
		var ps []Product
		err := json.Unmarshal(body, &ps)
		if err == nil {
			ctr.repo.Delete(ps)
			ctx.JSON(200, map[string]string{
				"Response": "Delete successful",
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

func (ctr ProductService) QueryByConditions(ctx *gin.Context) {
	min, max := getPriceRange(ctx)
	product := Product{
		Product_title: "%" + ctx.DefaultQuery("title", "%") + "%",
		Publisher:     "%" + ctx.DefaultQuery("publisher", "%") + "%",
	}
	ctx.JSON(200, ctr.repo.QueryByConditions(min, max, product))
}

func (serv ProductService) MaxPrice(ctx *gin.Context) {
	ctx.JSON(200, map[string]int{"MaxProductPrice": serv.repo.MaxPrice()})
}

func (serv ProductService) QueryNewest(ctx *gin.Context) {
	ran, err := strconv.Atoi(ctx.Param("range"))
	if err == nil {
		ctx.JSON(200, serv.repo.QueryNewest(ran))
	} else {
		ctx.JSON(400, map[string]string{
			"Response": "range must be integer",
		})
	}
}
