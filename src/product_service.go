package main

import (
	"context"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"
)

func (p Product) QueryAll(db *sqlx.DB, limit int64) ([]Product, error) {
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

func (p Product) Insert(db *sqlx.DB, ps []Product) (int64, error) {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	sqlC := SqlStringBuilder[Product]{
		Data:      ps,
		Form:      p,
		TableName: "product",
		Ids:       []string{"Product_id"},
	}
	stmt, err := db.PrepareContext(ctx, sqlC.GetInsertSQLString())

	if err != nil {
		panic(err)
	}

	params := func() []interface{} {
		var params []interface{}
		for _, v := range ps {
			if err == nil {
				params = append(params, v.Barcode, v.Publication_date, v.Product_title, v.Price, v.Publisher, v.Quantity, v.Description)
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
