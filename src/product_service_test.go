package main

import (
	"fmt"
	"testing"
)

func TestQuery(t *testing.T) {
	db, err := ConnectDB()

	if err == nil {
		var p Product
		products, err := p.Query(db, 50)

		if err == nil {
			fmt.Println(products)
		} else {
			t.Fatal(err)
		}
	} else {
		t.Fatal(err)
	}
}
