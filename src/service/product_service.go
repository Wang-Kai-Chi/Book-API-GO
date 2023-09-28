package service

import (
	"database/sql"

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

func (sqlS *ProductSqlStr) init() {
	sqlS.QueryWithLimit = GetSqlFromPath(sqlS.RelatedPath+sqlS.QueryWithLimit, sqlC)
	sqlS.QueryWithPriceRange = GetSqlFromPath(sqlS.RelatedPath+sqlS.QueryWithPriceRange, sqlC)
	sqlS.QueryByBarcode = GetSqlFromPath(sqlS.RelatedPath+sqlS.QueryByBarcode, sqlC)
	sqlS.Insert = GetSqlFromPath(sqlS.RelatedPath+sqlS.Insert, sqlC)
	sqlS.Update = GetSqlFromPath(sqlS.RelatedPath+sqlS.Update, sqlC)
	sqlS.Delete = GetSqlFromPath(sqlS.RelatedPath+sqlS.Delete, sqlC)
}

func NewProductSqlStr() *ProductSqlStr {
	return NewSqlS[*ProductSqlStr](`resource/sqlc/product/productSqlStr.json`, sqlC)
}

type ProductService struct{}

func (ser ProductService) QueryWithLimit(db *sqlx.DB, limit int64) []Product {
	return QueryEntity[Product](db, NewProductSqlStr().QueryWithLimit, limit)
}

func (ser ProductService) Insert(db *sqlx.DB, ps []Product) sql.Result {
	return ExecSql[[]Product](db, NewProductSqlStr().Insert, ps)
}

func (ser ProductService) QueryWithPriceRange(db *sqlx.DB, min int, max int) []Product {
	return QueryEntity[Product](db, NewProductSqlStr().QueryWithPriceRange, min, max)
}

func (ser ProductService) QueryByBarcode(db *sqlx.DB, code string) []Product {
	return QueryEntity[Product](db, NewProductSqlStr().QueryByBarcode, code)
}

func (ser ProductService) Update(db *sqlx.DB, ps []Product) sql.Result {
	return ExecSql[[]Product](db, NewProductSqlStr().Update, ps)
}

func (ser ProductService) Delete(db *sqlx.DB, ps []Product) sql.Result {
	return ExecSql[[]Product](db, NewProductSqlStr().Delete, ps)
}
