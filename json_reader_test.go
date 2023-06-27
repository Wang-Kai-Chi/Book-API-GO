package main

import (
	"testing"
)

func GetIsbnForTest() string {
	return "9789571313887"
}

func TestGetBookFromJson(t *testing.T) {
	isbn := GetBookFromJson("json/book_single.json").Isbn

	if isbn != GetIsbnForTest() {
		t.Fatal()
	}
}

func TestGetBookArrayFromJson(t *testing.T) {
	isbn := GetBookArrayFromJson("json/book_arr.json")[0].Isbn

	if isbn != GetIsbnForTest() {
		t.Fatal()
	}
}
