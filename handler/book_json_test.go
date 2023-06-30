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
	path := GetJsonFolder() + "book_single.json"
	isbn := GetBookFromJson(ReadFileAsString(path)).Isbn

	if isbn != GetIsbnForTest() {
		t.Fatal()
	}
}

func TestGetBookArrayFromJson(t *testing.T) {
	path := GetJsonFolder() + "book_arr.json"
	isbn := GetBookArrayFromJson(ReadFileAsString(path))[0].Isbn

	if isbn != GetIsbnForTest() {
		t.Fatal()
	}
}
