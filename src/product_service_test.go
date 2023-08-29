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

func TestQueryWithLimitProduct(t *testing.T) {
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

func convertCdToProducts() []Product {
	var cd Cd
	cds := cd.ConvertRaws(LoadData[[]RawCd]("../json/iknowbook.cd.json"))
	products := func() []Product {
		var ps []Product
		for _, v := range cds {
			ps = append(ps, v.Product_)
		}
		return ps
	}
	return products()
}

func TestPrintConvertedProduct(t *testing.T) {
	ps := convertCdToProducts()
	for i := 0; i < len(ps); i++ {
		if len(ps[i].Publication_date) == 0 {
			ps[i].Publication_date = "1975-01-01"
		} else {
			t.Log(ps[i].Publication_date)
		}
	}
	t.Log(len(ps))
}
func TestConvertAndInsertCds(t *testing.T) {
	ps := convertCdToProducts()

	for i := 0; i < len(ps); i++ {
		if len(ps[i].Publication_date) == 0 {
			ps[i].Publication_date = "1975-01-01"
		}
	}

	db, err := ConnectDB()

	if err == nil {
		var p Product
		res := p.Insert(db, ps)
		t.Log(res)
	} else {
		t.Fatal(err)
	}
}
