package main

import (
	"fmt"
	"testing"
)

func getProductForTest() []Product {
	return []Product{{
		Barcode:          "10000000",
		Product_title:    "testProduct",
		Publisher:        "testP",
		Publication_date: "1995-01-01",
		Price:            "69元",
		Quantity:         1,
		Description:      "none",
	}}
}

func TestQueryAllProduct(t *testing.T) {
	db, err := ConnectDB()

	if err == nil {
		var p Product
		products, err := p.QueryWithLimit(db, 50)

		if err == nil {
			fmt.Println(products)
		} else {
			t.Fatal(err)
		}
	} else {
		t.Fatal(err)
	}
}

func TestInsertProduct(t *testing.T) {
	db, err := ConnectDB()

	if err == nil {
		var p Product

		products := getProductForTest()
		res := p.Insert(db, products)
		t.Log(res)
	} else {
		t.Fatal(err)
	}
}
