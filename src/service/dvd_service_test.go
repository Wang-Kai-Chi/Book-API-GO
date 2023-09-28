package service

import (
	"testing"

	. "iknowbook.com/data"
)

func TestNewDvdSqlStr(t *testing.T) {
	sqlS := NewDvdSqlStr()
	t.Log(sqlS.RelatedPath)
	t.Log(sqlS.QueryWithLimit)
	t.Log(sqlS.Insert)
}

func TestQueryWithLimit(t *testing.T) {
	db, err := ConnectDB()
	if err == nil {
		serv := NewDvdService(db)
		dvds := serv.QueryWithLimit(50)
		t.Log(dvds)
	} else {
		t.Fatal(err)
	}
}

func TestInsertDvd(t *testing.T) {
	db, err := ConnectDB()
	if err == nil {
		serv := NewDvdService(db)
		dvds := []Dvd{
			{
				Product: Product{
					Barcode:       "4719851810640",
					Product_title: "夢土耳其三部曲-乳 DVD",
				},
			},
		}
		serv.Insert(dvds)
	} else {
		t.Fatal(err)
	}
}
