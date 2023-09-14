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
	QueryWithPriceRange string
	QueryByBarcode      string
	Insert              string
	Update              string
	Delete              string
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
	sqlS.QueryByBarcode = getSqlFromEmbededFolder(sqlS.RelatedPath + sqlS.QueryByBarcode)
	sqlS.Insert = getSqlFromEmbededFolder(sqlS.RelatedPath + sqlS.Insert)
	sqlS.Update = getSqlFromEmbededFolder(sqlS.RelatedPath + sqlS.Update)
	sqlS.Delete = getSqlFromEmbededFolder(sqlS.RelatedPath + sqlS.Delete)
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

func execSql(db *sqlx.DB, str string, ps []Product) sql.Result {
	res, err := db.NamedExec(str, ps)
	if err != nil {
		panic(err)
	}
	db.Close()

	return res
}

func (ser ProductService) QueryWithLimit(db *sqlx.DB, limit int64) []Product {
	return queryProducts(db, NewProductSqlStr().QueryWithLimit, limit)
}

func (ser ProductService) Insert(db *sqlx.DB, ps []Product) sql.Result {
	return execSql(db, NewProductSqlStr().Insert, ps)
}

func (ser ProductService) QueryWithPriceRange(db *sqlx.DB, min int, max int) []Product {
	return queryProducts(db, NewProductSqlStr().QueryWithPriceRange, min, max)
}

func (ser ProductService) QueryByBarcode(db *sqlx.DB, code string) []Product {
	return queryProducts(db, NewProductSqlStr().QueryByBarcode, code)
}

func (ser ProductService) Update(db *sqlx.DB, ps []Product) sql.Result {
	return execSql(db, NewProductSqlStr().Update, ps)
}

func (ser ProductService) Delete(db *sqlx.DB, ps []Product) sql.Result {
	return execSql(db, NewProductSqlStr().Delete, ps)
}
