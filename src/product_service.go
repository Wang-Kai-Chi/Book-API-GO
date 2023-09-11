package main

import (
	"database/sql"
	"embed"
	"encoding/json"

	"github.com/jmoiron/sqlx"
)

type ProductSqlStr struct {
	RelatedPath         string
	QueryWithLimit      string
	Insert              string
	QueryWithPriceRange string
	QueryByBarcode      string
}

//go:embed resource/sqlc/product/*
var productSqlC embed.FS

func initProductSql(sqlS *ProductSqlStr) {
	getSqlFromEmbededFolder := func(path string) string {
		data, _ := productSqlC.ReadFile(path)
		return string(data)
	}
	sqlS.QueryWithLimit = getSqlFromEmbededFolder(sqlS.RelatedPath + sqlS.QueryWithLimit)

	sqlS.QueryWithPriceRange = getSqlFromEmbededFolder(sqlS.RelatedPath + sqlS.QueryWithPriceRange)

	sqlS.Insert = getSqlFromEmbededFolder(sqlS.RelatedPath + sqlS.Insert)

	sqlS.QueryByBarcode = getSqlFromEmbededFolder(sqlS.RelatedPath + sqlS.QueryByBarcode)
}

func NewProductSqlStr() ProductSqlStr {
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

func queryProducts(db *sqlx.DB, sqlStr string, params ...interface{}) []Product {
	mustGetRowsFromQuery := func(db *sqlx.DB, sqlStr string, params ...interface{}) *sqlx.Rows {
		query, args, err := sqlx.In(sqlStr, params...)
		if err != nil {
			panic(err)
		}
		query = db.Rebind(query)
		rows, err := db.Queryx(query, args...)
		if err != nil {
			panic(err)
		}
		return rows
	}
	mustGetProductFromRows := func(rows *sqlx.Rows) []Product {
		var products []Product
		for rows.Next() {
			var p Product
			err := rows.StructScan(&p)
			if err != nil {
				panic(err)
			}
			products = append(products, p)
		}
		return products
	}
	rows := mustGetRowsFromQuery(db, sqlStr, params...)
	products := mustGetProductFromRows(rows)
	db.Close()
	return products
}

func (service ProductService) QueryWithLimit(db *sqlx.DB, limit int64) []Product {
	return queryProducts(db, NewProductSqlStr().QueryWithLimit, limit)
}

func (service ProductService) Insert(db *sqlx.DB, ps []Product) sql.Result {
	sqlC := NewProductSqlStr()
	sqlStr := sqlC.Insert
	res, err := db.NamedExec(sqlStr, ps)
	if err != nil {
		panic(err)
	}
	db.Close()

	return res
}

func (service ProductService) QueryWithPriceRange(db *sqlx.DB, min int, max int) []Product {
	return queryProducts(db, NewProductSqlStr().QueryWithPriceRange, min, max)
}

func (service ProductService) QueryByBarcode(db *sqlx.DB, code string) []Product {
	return queryProducts(db, NewProductSqlStr().QueryByBarcode, code)
}
