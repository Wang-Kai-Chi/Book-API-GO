package main

import (
	"fmt"
	"testing"
)

func isbnSample() string {
	return "9789571313887"
}

func upacSample() string {
	return "4715219794386"
}

func TestMustGetDataFromJson(t *testing.T) {
	loader := NewDataTestLoader()

	rawBook := loader.getSingleRawBook

	if rawBook.Isbn == isbnSample() {
		fmt.Println(rawBook)
	} else {
		fmt.Println(rawBook)
		t.Fatal()
	}

	book := loader.getSingleBook

	if book.Product_.Barcode == isbnSample() {
		fmt.Println(book)
	} else {
		t.Fatal()
	}

	rawBooks := loader.getRawBooks

	if rawBooks[0].Isbn == isbnSample() {
		fmt.Println(rawBooks)
	} else {
		t.Fatal()
	}

	dvd := loader.getSingleRawDvd

	if dvd.Barcode == upacSample() {
		fmt.Println(dvd)
	} else {
		fmt.Println(dvd)
		t.Fatal()
	}

	dvds := loader.getRawDvds

	if dvds[0].Barcode == upacSample() {
		fmt.Println(dvds)
	} else {
		t.Fatal()
	}
}
