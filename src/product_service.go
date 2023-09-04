package main

import (
	"database/sql"
	"encoding/json"

	"github.com/jmoiron/sqlx"
	"iknowbook.com/handler"
)

type ProductSqlStr struct {
	RelatedPath         string
	QueryWithLimit      string
	Insert              string
	QueryWithPriceRange string
}

func NewProductSqlStr() ProductSqlStr {
	var sqlS ProductSqlStr

	err := json.Unmarshal(handler.MustReadFile("./resource/sqlc/product/productSqlStr.json"), &sqlS)

	if err == nil {
		sqlS.QueryWithLimit = handler.ReadFileAsString(sqlS.RelatedPath + sqlS.QueryWithLimit)
		sqlS.QueryWithPriceRange = handler.ReadFileAsString(sqlS.RelatedPath + sqlS.QueryWithPriceRange)
		sqlS.Insert = handler.ReadFileAsString(sqlS.RelatedPath + sqlS.Insert)
	} else {
		panic(err)
	}

	return sqlS
}

type ProductService struct{}

func MustGetRowsFromQuery(db *sqlx.DB, sqlStr string, params ...interface{}) *sqlx.Rows {
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

func MustGetProductFromRows(rows *sqlx.Rows) []Product {
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

func (service ProductService) QueryWithLimit(db *sqlx.DB, limit int64) []Product {
	sqlc := NewProductSqlStr()

	rows := MustGetRowsFromQuery(db, sqlc.QueryWithLimit, limit)

	products := MustGetProductFromRows(rows)

	db.Close()

	return products
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
	sqlc := NewProductSqlStr()

	rows := MustGetRowsFromQuery(db, sqlc.QueryWithPriceRange, min, max)

	products := MustGetProductFromRows(rows)

	db.Close()

	return products
}
