package cd

import (
	"testing"

	. "iknowbook.com/data"
	. "iknowbook.com/db"
)

func TestQueryCdWithLimit(t *testing.T) {
	db, err := ConnectDB()

	if err == nil {
		serv := NewCdRepository(db)
		cds := serv.QueryWithLimit(50)
		t.Log(cds)
	} else {
		t.Fatal(err)
	}
}

func TestInsertCds(t *testing.T) {
	db, err := ConnectDB()

	if err == nil {
		serv := NewCdRepository(db)
		cds := []Cd{
			{Product: Product{
				Barcode:       "028947757092",
				Product_title: "數位古典大師-巴哈",
			}},
		}
		t.Log(serv.Insert(cds).RowsAffected())
	} else {
		t.Fatal(err)
	}
}
