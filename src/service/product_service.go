package service

import (
	"database/sql"
	"embed"
	"encoding/json"

	. "iknowbook.com/data"

	"github.com/jmoiron/sqlx"
)

type ProductSqlStr struct {
	RelatedPath         string
	QueryWithLimit      string
	QueryWithPriceRange string
	QueryByBarcode      string
	Insert              string
	Update              string
	Delete              string
}

//go:embed resource/sqlc/product/*
var productSqlC embed.FS

func NewProductSqlStr() ProductSqlStr {
	initProductSql := func(sqlS *ProductSqlStr) {
		getSqlFromEmbededFolder := func(path string) string {
			data, _ := productSqlC.ReadFile(path)
			return string(data)
		}
		sqlS.QueryWithLimit = getSqlFromEmbededFolder(sqlS.RelatedPath + sqlS.QueryWithLimit)
		sqlS.QueryWithPriceRange = getSqlFromEmbededFolder(sqlS.RelatedPath + sqlS.QueryWithPriceRange)
		sqlS.QueryByBarcode = getSqlFromEmbededFolder(sqlS.RelatedPath + sqlS.QueryByBarcode)
		sqlS.Insert = getSqlFromEmbededFolder(sqlS.RelatedPath + sqlS.Insert)
		sqlS.Update = getSqlFromEmbededFolder(sqlS.RelatedPath + sqlS.Update)
		sqlS.Delete = getSqlFromEmbededFolder(sqlS.RelatedPath + sqlS.Delete)
	}
	data, err := productSqlC.ReadFile("resource/sqlc/product/productSqlStr.json")
	var sqlS ProductSqlStr
	if err != nil {
		panic(err)
	} else {
		err := json.Unmarshal(data, &sqlS)
		if err != nil {
			panic(err)
		}
		initProductSql(&sqlS)
	}
	return sqlS
}

type ProductService struct{}

func (ser ProductService) QueryWithLimit(db *sqlx.DB, limit int64) []Product {
	return QueryEntity[Product](db, NewProductSqlStr().QueryWithLimit, limit)
}

func (ser ProductService) Insert(db *sqlx.DB, ps []Product) sql.Result {
	return ExecSql[Product](db, NewProductSqlStr().Insert, ps)
}

func (ser ProductService) QueryWithPriceRange(db *sqlx.DB, min int, max int) []Product {
	return QueryEntity[Product](db, NewProductSqlStr().QueryWithPriceRange, min, max)
}

func (ser ProductService) QueryByBarcode(db *sqlx.DB, code string) []Product {
	return QueryEntity[Product](db, NewProductSqlStr().QueryByBarcode, code)
}

func (ser ProductService) Update(db *sqlx.DB, ps []Product) sql.Result {
	return ExecSql[Product](db, NewProductSqlStr().Update, ps)
}

func (ser ProductService) Delete(db *sqlx.DB, ps []Product) sql.Result {
	return ExecSql[Product](db, NewProductSqlStr().Delete, ps)
}
