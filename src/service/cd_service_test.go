package service

import (
	"testing"

	. "iknowbook.com/data"
)

func TestNewCdSqlStr(t *testing.T) {
	sqlS := NewCdSqlStr()
	t.Log(sqlS.QueryWithLimit)
	t.Log(sqlS.Insert)
}

func TestQueryCdWithLimit(t *testing.T) {
	db, err := ConnectDB()

	if err == nil {
		serv := NewCdService(db)
		cds := serv.QueryWithLimit(50)
		t.Log(cds)
	} else {
		t.Fatal(err)
	}
}

func TestInsertCds(t *testing.T) {
	db, err := ConnectDB()

	if err == nil {
		serv := NewCdService(db)
		cds := []Cd{
			{Product: Product{
				Barcode:       "028947757092",
				Product_title: "數位古典大師-巴哈",
			}},
		}
		serv.Insert(cds)
	} else {
		t.Fatal(err)
	}
}
