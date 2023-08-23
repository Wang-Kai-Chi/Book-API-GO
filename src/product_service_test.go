package main

import (
	"fmt"
	"reflect"
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

func TestQueryProduct(t *testing.T) {
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

func TestReflect(t *testing.T) {
	products := GetProductForTest()

	t.Log(reflect.TypeOf(products[0]))
	t.Log(reflect.ValueOf(products[0]))

	elm := reflect.ValueOf(&products[0]).Elem()
	t.Log(elm.Type().Field(0).Name)
}
