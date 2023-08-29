package main

import (
	"database/sql"
	"strconv"

	"github.com/jmoiron/sqlx"
)

func (p Product) QueryWithLimit(db *sqlx.DB, limit int64) ([]Product, error) {
	rows, err := db.Queryx("SELECT * FROM product LIMIT " + strconv.FormatInt(limit, 10))

	var products []Product

	for rows.Next() {
		var p Product

		err := rows.StructScan(&p)
		if err != nil {
			panic(err)
		}
		products = append(products, p)
	}
	db.Close()
	return products, err
}

func (p Product) Insert(db *sqlx.DB, ps []Product) sql.Result {
	sqlStr := NewProductSqlC().GetInsertString()
	res, err := db.NamedExec(sqlStr, ps)
	if err != nil {
		panic(err)
	}
	db.Close()

	return res
}
