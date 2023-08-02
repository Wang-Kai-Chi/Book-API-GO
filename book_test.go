package main

import (
	"encoding/json"
	"testing"

	"iknowbook.com/handler"
)

func isbnSample() string {
	return "9789571313887"
}

func mustGetBookFromJson(content string) Book {
	var book Book
	err := json.Unmarshal([]byte(content), &book)
	if err != nil {
		panic(err)
	}
	return book
}

func mustGetBookArrayFromJson(content string) []Book {
	var books []Book
	err := json.Unmarshal([]byte(content), &books)
	if err != nil {
		panic(err)
	}
	return books
}

func getSingleBook() Book {
	return mustGetBookFromJson(handler.ReadFileAsString("../json/book_single.json"))
}

func getBooks() []Book {
	return mustGetBookArrayFromJson(handler.ReadFileAsString("../json/book_arr.json"))
}

func TestMustGetBookFromJson(t *testing.T) {
	isbn := getSingleBook().Isbn
	if isbn != isbnSample() {
		t.Fatal()
	}
}

func TestMustGetBookArrayFromJson(t *testing.T) {
	isbn := getBooks()[0].Isbn
	if isbn != isbnSample() {
		t.Fatal()
	}
}
