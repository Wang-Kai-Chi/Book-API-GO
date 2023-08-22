package main

import (
	"testing"
)

func TestConnectDB(t *testing.T) {
	_, err := ConnectDB()

	if err != nil {
		t.Fatal()
	}
}
