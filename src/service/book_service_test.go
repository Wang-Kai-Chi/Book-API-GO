package service

import (
	"testing"

	. "iknowbook.com/data"
)

func TestNewBookSqlStr(t *testing.T) {
	sqlS := NewBookSqlStr()
	t.Log(sqlS.QueryByLimit)
	t.Log(sqlS.Insert)
}

func TestQueryByLimit(t *testing.T) {
	db, err := ConnectDB()
	if err == nil {
		serv := NewBookService(db)
		books := serv.QueryByLimit(50)
		t.Log(books)
	} else {
		t.Fatal(err)
	}
}

func TestInsert(t *testing.T) {
	db, err := ConnectDB()
	if err == nil {
		serv := NewBookService(db)
		books := []Book{
			{
				Product: Product{
					Barcode:       "9789863241195",
					Product_title: "妖怪少爺 (23)",
				},
				Author:     "",
				Translator: "",
				Language:   "中文",
				Category:   "漫畫",
			},
			{
				Product: Product{
					Barcode:       "9789863241195",
					Product_title: "妖怪少爺 (23)",
				},
				Author:     "",
				Translator: "",
				Language:   "中文",
				Category:   "漫畫",
			},
		}
		serv.Insert(books)
	} else {
		t.Fatal(err)
	}
}
