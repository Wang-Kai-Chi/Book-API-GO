package handler

import (
	"testing"
)

func GetIsbnForTest() string {
	return "9789571313887"
}

func GetJsonFolder() string {
	return "../json/"
}

func TestGetBookFromJson(t *testing.T) {
	isbn := GetBookFromJson(GetJsonFolder() + "book_single.json").Isbn

	if isbn != GetIsbnForTest() {
		t.Fatal()
	}
}

func TestGetBookArrayFromJson(t *testing.T) {
	isbn := GetBookArrayFromJson(GetJsonFolder() + "book_arr.json")[0].Isbn

	if isbn != GetIsbnForTest() {
		t.Fatal()
	}
}
