package product

import (
	"fmt"
	"strconv"
	"testing"

	. "iknowbook.com/app/data"
	. "iknowbook.com/app/db"
)

func getProductForTest() []Product {
	return []Product{{
		Product_id:       1464,
		Barcode:          "9789861043128",
		Product_title:    "妖怪少爺 (6)",
		Publisher:        "東立",
		Publication_date: "2022-03-19T00:00:00Z",
		Price:            "85元",
		Description:      "none",
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
		products := p.QueryByBarcode("9789861043128")
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
		num, err := res.RowsAffected()
		if err == nil {
			rowCount := strconv.Itoa(int(num))
			if err == nil {
				t.Log("Update success, " + rowCount + " affected")
			}
		}
	} else {
		t.Fatal(err)
	}
}

func TestDelete(t *testing.T) {
	db, err := ConnectDB()
	if err == nil {
		p := NewProductRepository(db)
		res := p.Delete([]Product{{
			Product_id: 2066,
		}})
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

func TestQueryNewest(t *testing.T) {
	db, err := ConnectDB()
	if err == nil {
		repo := NewProductRepository(db)
		ps := repo.QueryNewest(10)
		t.Log(ps)
	} else {
		t.Fatal(err)
	}
}
