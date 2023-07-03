package handler

import (
	"fmt"
	"testing"
)

const ISBN = "9789571313887"
const FOLDER = "../json/"

func GetSingleBook() Book {
	return GetBookFromJson(ReadFileAsString("../json/book_single.json"))
}

func GetBooks() []Book {
	return GetBookArrayFromJson(ReadFileAsString(FOLDER + "book_arr.json"))
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

func TestBookToString(t *testing.T) {
	fmt.Println(BookToString(GetSingleBook()))
}

func TestBooksToString(t *testing.T) {
	fmt.Println(BooksToString(GetBooks()))
}
