package handler

import (
	"testing"
)

const ISBN = "9789571313887"
const FOLDER = "../json/"

func TestGetBookFromJson(t *testing.T) {
	path := FOLDER + "book_single.json"
	isbn := GetBookFromJson(ReadFileAsString(path)).Isbn

	if isbn != ISBN {
		t.Fatal()
	}
}

func TestGetBookArrayFromJson(t *testing.T) {
	path := FOLDER + "book_arr.json"
	isbn := GetBookArrayFromJson(ReadFileAsString(path))[0].Isbn

	if isbn != ISBN {
		t.Fatal()
	}
}
