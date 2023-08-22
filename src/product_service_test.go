package main

import (
	"fmt"
	"testing"
)

func TestQuery(t *testing.T) {
	db, err := ConnectDB()

	if err == nil {
		var p Product
		products, err := p.Query(db, 50)

		if err == nil {
			fmt.Println(products)
		} else {
			t.Fatal(err)
		}
	} else {
		t.Fatal(err)
	}
}

func TestInsert(t *testing.T) {
	db, err := ConnectDB()

	if err == nil {
		var p Product

		products := []Product{{
			Barcode:         "10000000",
			Product_title:   "testProduct",
			Publisher:       "testP",
			PublicationDate: "1995-01-01",
			Price:           "69元",
			Quantity:        1,
		}}
		rows, err := p.Insert(db, products)

		if err == nil {
			fmt.Println(rows)
		} else {
			t.Fatal(err)
		}
	} else {
		t.Fatal(err)
	}
}
