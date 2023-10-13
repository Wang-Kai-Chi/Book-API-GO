package dvd

import (
	"testing"

	. "iknowbook.com/data"
	. "iknowbook.com/db"
)

func TestQueryWithLimit(t *testing.T) {
	db, err := ConnectDB()
	if err == nil {
		serv := NewDvdRepository(db)
		dvds := serv.QueryWithLimit(50)
		t.Log(dvds)
	} else {
		t.Fatal(err)
	}
}

func TestInsertDvd(t *testing.T) {
	db, err := ConnectDB()
	if err == nil {
		serv := NewDvdRepository(db)
		dvds := []Dvd{
			{
				Product: Product{
					Barcode:       "4719851810640",
					Product_title: "夢土耳其三部曲-乳 DVD",
				},
			},
		}
		t.Log(serv.Insert(dvds).RowsAffected())
	} else {
		t.Fatal(err)
	}
}
