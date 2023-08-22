package main

import (
	"context"
	"database/sql"
	"strconv"
	"strings"
	"time"
)

func (p Product) Query(db *sql.DB, limit int) ([]Product, error) {
	rows, err := db.Query(
		"SELECT product_id, barcode, product_title, publisher, publication_date, price, quantity, description FROM product LIMIT " + string(rune(limit)))

	var products []Product

	for rows.Next() {
		var p Product

		err := rows.Scan(
			&p.Id,
			&p.Barcode,
			&p.Product_title,
			&p.Publisher,
			&p.PublicationDate,
			&p.Price,
			&p.Quantity,
			&p.Description,
		)
		if err == nil {
			products = append(products, p)
		}
	}
	db.Close()

	return products, err
}

func (p Product) Insert(db *sql.DB, ps []Product) (int64, error) {
	sqlStr := "INSERT INTO product(barcode, product_title, publisher, publication_date, price, quantity) VALUES "

	var inserts []string
	var params []interface{}

	for _, v := range ps {
		inserts = append(inserts, "($1, $2, $3, $4, $5, $6)")

		price, err := strconv.Atoi(strings.ReplaceAll(v.Price, "元", ""))

		if err == nil {
			params = append(params, v.Barcode, v.Product_title, v.Publisher, v.PublicationDate, price, v.Quantity)
		}
	}

	insertVals := strings.Join(inserts, ",")
	sqlStr = sqlStr + insertVals

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	stmt, err := db.PrepareContext(ctx, sqlStr)

	if err != nil {
		panic(err)
	}

	res, err := stmt.ExecContext(ctx, params...)
	if err != nil {
		panic(err)
	}
	rows, err := res.RowsAffected()

	return rows, err
}
