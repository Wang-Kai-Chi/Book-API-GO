package product

import (
	"fmt"
	"testing"

	. "iknowbook.com/data"
	. "iknowbook.com/db"
)

func getProductForTest() []Product {
	return []Product{{
		Barcode:          "1111",
		Product_title:    "test1111",
		Publisher:        "test",
		Publication_date: "1995-01-01",
		Price:            "69",
		Description:      "",
		Quantity:         1,
	}}
}

func TestQueryWithLimitProduct(t *testing.T) {
	db, err := ConnectDB()

	if err == nil {
		p := NewProductRepository(db)
		products := p.QueryWithLimit(50)
		fmt.Println(products)
	} else {
		t.Fatal(err)
	}
}

func TestInsertProduct(t *testing.T) {
	db, err := ConnectDB()

	if err == nil {
		p := NewProductRepository(db)

		products := getProductForTest()
		res := p.Insert(products)
		t.Log(res)
	} else {
		t.Fatal(err)
	}
}
func TestQueryWithPriceRange(t *testing.T) {
	db, err := ConnectDB()

	if err == nil {
		p := NewProductRepository(db)

		products := p.QueryWithPriceRange(50, 100)
		t.Log(products)
	} else {
		t.Fatal(err)
	}
}

func TestQueryByBarcode(t *testing.T) {
	db, err := ConnectDB()
	if err == nil {
		p := NewProductRepository(db)
		products := p.QueryByBarcode("602508588662")
		t.Log(products)
	} else {
		t.Fatal(err)
	}
}

func TestUpdate(t *testing.T) {
	db, err := ConnectDB()
	if err == nil {
		p := NewProductRepository(db)
		res := p.Update(getProductForTest())
		t.Log(res)
	} else {
		t.Fatal(err)
	}
}

func TestDelete(t *testing.T) {
	db, err := ConnectDB()
	if err == nil {
		p := NewProductRepository(db)
		res := p.Delete(getProductForTest())
		t.Log(res)
	} else {
		t.Fatal(err)
	}
}

func TestQueryByConditions(t *testing.T) {
	db, err := ConnectDB()
	if err == nil {
		serv := NewProductRepository(db)
		cond := Product{
			Product_title: "妖怪%",
			Publisher:     "%",
		}
		ps := serv.QueryByConditions(0, 250, cond)
		t.Log(ps)
	} else {
		t.Fatal(err)
	}
}

func TestMaxPrice(t *testing.T) {
	db, err := ConnectDB()
	if err == nil {
		repo := NewProductRepository(db)
		max := repo.MaxPrice()
		t.Log(max)
	} else {
		t.Fatal(err)
	}
}
