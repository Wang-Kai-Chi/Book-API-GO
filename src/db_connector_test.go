package main

import (
	"fmt"
	"testing"
)

func TestConnectDB(t *testing.T) {
	_, err := ConnectDB()

	if err != nil {
		t.Fatal()
	}
}

func TestQueryProduct(t *testing.T) {
	db, err := ConnectDB()
	if err == nil {
		products, err := QueryProduct(db)
		if err != nil {
			t.Log(err)
			t.Fatal()
		} else {
			fmt.Println(products)
		}
	} else {
		t.Fatal()
	}
}
