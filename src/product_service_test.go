package main

import (
	"fmt"
	"testing"

	. "iknowbook.com/data"
)

func getProductForTest() []Product {
	return []Product{{
		Product_id:       1982,
		Barcode:          "10000000",
		Product_title:    "testProduct",
		Publisher:        "testPpp",
		Publication_date: "1995-01-01",
		Price:            "69元",
		Quantity:         1,
		Description:      "none",
	}}
}

func TestNewProductSqlStr(t *testing.T) {
	sqlS := NewProductSqlStr()
	t.Log(sqlS.RelatedPath)
	t.Log(sqlS.QueryWithLimit)
	t.Log(sqlS.Insert)
	t.Log(sqlS.QueryWithPriceRange)
	t.Log(sqlS.QueryByBarcode)
	t.Log(sqlS.Delete)
	t.Log(sqlS.Update)
}

func TestQueryWithLimitProduct(t *testing.T) {
	db, err := ConnectDB()

	if err == nil {
		var p ProductService
		products := p.QueryWithLimit(db, 50)
		fmt.Println(products)
	} else {
		t.Fatal(err)
	}
}

func TestInsertProduct(t *testing.T) {
	db, err := ConnectDB()

	if err == nil {
		var p ProductService

		products := getProductForTest()
		res := p.Insert(db, products)
		t.Log(res)
	} else {
		t.Fatal(err)
	}
}
func TestQueryWithPriceRange(t *testing.T) {
	db, err := ConnectDB()

	if err == nil {
		var p ProductService

		products := p.QueryWithPriceRange(db, 50, 100)
		t.Log(products)
	} else {
		t.Fatal(err)
	}
}

func TestQueryByBarcode(t *testing.T) {
	db, err := ConnectDB()
	if err == nil {
		var p ProductService
		products := p.QueryByBarcode(db, "602508588662")
		t.Log(products)
	} else {
		t.Fatal(err)
	}
}

func TestUpdate(t *testing.T) {
	db, err := ConnectDB()
	if err == nil {
		var p ProductService
		res := p.Update(db, getProductForTest())
		t.Log(res)
	} else {
		t.Fatal(err)
	}
}

func TestDelete(t *testing.T) {
	db, err := ConnectDB()
	if err == nil {
		var p ProductService
		res := p.Delete(db, getProductForTest())
		t.Log(res)
	} else {
		t.Fatal(err)
	}
}
