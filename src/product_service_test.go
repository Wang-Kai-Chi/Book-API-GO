package main

import (
	"fmt"
	"testing"
)

func GetProductForTest() []Product {
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
		products, err := p.QueryAll(db, 50)

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

		products := GetProductForTest()
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

func TestGetInsertSQLString(t *testing.T) {
	var p Product
	s := GetInsertSQLString[Product](GetProductForTest(), p, "product", []string{"Product_id"})
	t.Log(s)
}
