package main

import "database/sql"

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
