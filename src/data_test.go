package main

import (
	"encoding/json"
	"testing"

	"iknowbook.com/handler"
)

func TestMustGetDataFromJson(t *testing.T) {
	if getSingleBook().Isbn != isbnSample() {
		t.Fatal()
	}

	if getBooks()[0].Isbn != isbnSample() {
		t.Fatal()
	}

	if getSingleDvd().Barcode != upacSample() {
		t.Fatal()
	}

	if getDvds()[0].Barcode != upacSample() {
		t.Fatal()
	}
}

func mustGetDataFromJson[T Data](content string) T {
	var entity T
	err := json.Unmarshal([]byte(content), &entity)
	if err != nil {
		panic(err)
	}
	return entity
}

func getSingleBook() Book {
	path := "../json/book_single.json"
	return mustGetDataFromJson[Book](handler.ReadFileAsString(path))
}

func getBooks() []Book {
	path := "../json/book_arr.json"
	return mustGetDataFromJson[[]Book](handler.ReadFileAsString(path))
}

func getSingleDvd() Dvd {
	path := "../json/dvd_single.json"
	return mustGetDataFromJson[Dvd](handler.ReadFileAsString(path))
}

func getDvds() []Dvd {
	path := "../json/dvd_array.json"
	return mustGetDataFromJson[[]Dvd](handler.ReadFileAsString(path))
}

func isbnSample() string {
	return "9789571313887"
}

func upacSample() string {
	return "4715219794386"
}
