package main

import (
	"context"
	"database/sql"
	"strconv"
	"strings"
	"time"
)

func (p Product) Query(db *sql.DB, limit int64) ([]Product, error) {
	rows, err := db.Query(
		"SELECT product_id, barcode, product_title, publisher, publication_date, price, quantity, description FROM product LIMIT " + strconv.FormatInt(limit, 10))

	var products []Product

	for rows.Next() {
		var p Product

		err := rows.Scan(
			&p.Product_id,
			&p.Barcode,
			&p.Product_title,
			&p.Publisher,
			&p.Publication_date,
			&p.Price,
			&p.Quantity,
			&p.Description,
		)
		if err != nil {
			panic(err)
		}
		products = append(products, p)
	}
	db.Close()

	return products, err
}

func (p Product) Insert(db *sql.DB, ps []Product) (int64, error) {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	sqlStr := func() string {
		sqlStr := "INSERT INTO product(barcode, product_title, publisher, publication_date, price, quantity, description) VALUES "

		var inserts []string
		for i := 0; i < len(ps); i++ {
			inserts = append(inserts, "($1, $2, $3, $4, $5, $6, $7)")
		}
		insertVals := strings.Join(inserts, ",")
		sqlStr = sqlStr + insertVals
		return sqlStr
	}
	stmt, err := db.PrepareContext(ctx, sqlStr())

	if err != nil {
		panic(err)
	}

	params := func() []interface{} {
		var params []interface{}
		for _, v := range ps {
			price, err := strconv.Atoi(strings.ReplaceAll(v.Price, "元", ""))

			if err == nil {
				params = append(params, v.Barcode, v.Product_title, v.Publisher, v.Publication_date, price, v.Quantity, v.Description)
			}
		}
		return params
	}
	res, err := stmt.ExecContext(ctx, params()...)
	if err != nil {
		panic(err)
	}
	rows, err := res.RowsAffected()

	return rows, err
}
