package main

import (
	"testing"
)

func TestConnectDB(t *testing.T) {
	db, err := ConnectDB()

	if err != nil {
		t.Fatal()
	}
	db.Close()
}
