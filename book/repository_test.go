package book

import (
	"testing"

	. "iknowbook.com/data"
	. "iknowbook.com/db"
)

func TestQueryByLimit(t *testing.T) {
	db, err := ConnectDB()
	if err == nil {
		serv := NewBookRepository(db)
		books := serv.QueryWithLimit(50)
		t.Log(books)
	} else {
		t.Fatal(err)
	}
}

func TestInsert(t *testing.T) {
	db, err := ConnectDB()
	if err == nil {
		serv := NewBookRepository(db)
		books := []Book{
			{
				Product: Product{
					Barcode:          "1111",
					Product_title:    "test1111",
					Publisher:        "test",
					Publication_date: "1995-01-01",
					Price:            "69",
					Description:      "",
					Quantity:         1,
				},
				Author:     "",
				Translator: "",
				Language:   "",
				Category:   "",
			},
		}
		t.Log(serv.Insert(books).LastInsertId())
	} else {
		t.Fatal(err)
	}
}

func TestQueryBookByConditions(t *testing.T) {
	db, err := ConnectDB()
	if err == nil {
		serv := NewBookRepository(db)
		cond := Book{
			Product: Product{
				Product_title: "%",
				Publisher:     "%",
			},
			Author:     "%羅%",
			Translator: "%",
			Language:   "%",
			Category:   "%",
		}
		books := serv.QueryByConditions(0, 600, cond)
		t.Log(books)
	} else {
		t.Fatal(err)
	}

}

func TestBookUpdate(t *testing.T) {
	db, err := ConnectDB()
	if err == nil {
		serv := NewBookRepository(db)
		books := []Book{
			{Product: Product{
				Product_id:       1857,
				Barcode:          "9789863371908",
				Product_title:    "鬼灯的冷徹(10)",
				Publisher:        "東立",
				Publication_date: "2022-03-19",
				Quantity:         1,
				Description:      "none",
				Price:            "110元",
			},
				Author:     "江口夏実",
				Translator: "鄭啟旭",
				Language:   "繁體中文",
				Category:   "漫畫",
			},
		}
		t.Log(serv.Update(books).RowsAffected())
	} else {
		t.Fatal(err)
	}
}
