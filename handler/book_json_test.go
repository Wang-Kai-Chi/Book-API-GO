package handler

import (
	"testing"
)

const (
	ISBN = "9789571313887"
)

func GetSingleBook() Book {
	return MustGetBookFromJson(ReadFileAsString("../json/book_single.json"))
}

func GetBooks() []Book {
	return MustGetBookArrayFromJson(ReadFileAsString("../json/book_arr.json"))
}

func TestGetBookFromJson(t *testing.T) {
	isbn := GetSingleBook().Isbn
	if isbn != ISBN {
		t.Fatal()
	}
}

func TestGetBookArrayFromJson(t *testing.T) {
	isbn := GetBooks()[0].Isbn
	if isbn != ISBN {
		t.Fatal()
	}
}
