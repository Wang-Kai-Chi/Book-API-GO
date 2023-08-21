package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	config, err := GetConfig()

	var dataBase *sql.DB
	var newErr error

	if err == nil {
		db, err := sql.Open(config.DriverName, config.DBConnection)
		dataBase = db
		newErr = err
	}

	return dataBase, newErr
}

func QueryProduct(db *sql.DB) ([]Product, error) {
	rows, err := db.Query(
		"SELECT product_id, barcode, product_title, publisher, publication_date, price, quantity, description FROM product LIMIT 50")

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
