package main

import (
	"fmt"
	"testing"
)

func TestGetConfig(t *testing.T) {
	con, err := GetConfig()

	if err == nil {
		fmt.Println(con)
	} else {
		t.Log(err)
		t.Fatal()
	}
}
