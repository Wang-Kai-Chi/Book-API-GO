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
	out := products()
	initNullPublicaionDate(out)
	return out
}
func TestPrintConvertedProduct(t *testing.T) {
	ps := convertCdToProducts()
	for _, v := range ps {
		t.Log(v.Publication_date)
	}
}

func convertDvdToProducts() []Product {
	var dvd Dvd
	dvds := dvd.ConvertRaws(LoadData[[]RawDvd]("../json/iknowbook.dvd.json"))
	ps := func() []Product {
		var temp []Product
		for _, v := range dvds {
			temp = append(temp, v.Product_)
		}
		return temp
	}
	out := ps()
	initNullPublicaionDate(out)
	return out
}

func TestConvertDvdToProducts(t *testing.T) {
	ps := convertDvdToProducts()
	for i := 0; i < 50; i++ {
		t.Log(ps[i])
	}
}

func convertBookToProduct() []Product {
	var book Book
	books := book.ConvertRaws(LoadData[[]RawBook]("../json/iknowbook.book.json"))
	ps := func() []Product {
		var temp []Product
		for _, v := range books {
			temp = append(temp, v.Product_)
		}
		return temp
	}
	out := ps()
	initNullPublicaionDate(out)
	return out
}

func TestConvertBookToProduct(t *testing.T) {
	ps := convertBookToProduct()
	for i := 0; i < 50; i++ {
		t.Log(ps[i])
	}
}

func initNullPublicaionDate(ps []Product) {
	for i := 0; i < len(ps); i++ {
		if len(ps[i].Publication_date) == 0 {
			ps[i].Publication_date = "1975-01-01"
		}
	}
}
func TestConvertAndInsertProducts(t *testing.T) {
	ps := convertBookToProduct()
	db, err := ConnectDB()

	if err == nil {
		var p Product
		res := p.Insert(db, ps)
		t.Log(res)
	} else {
		t.Fatal(err)
	}
}
