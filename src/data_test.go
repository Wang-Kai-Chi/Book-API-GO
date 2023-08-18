package main

import (
	"testing"
)

func isbnSample() string {
	return "9789571313887"
}

func upacSample() string {
	return "4715219794386"
}

func isIsbnNotEquals(isbn string) bool {
	return isbn != isbnSample()
}

func isUpacNotEquals(upac string) bool {
	return upac != upacSample()
}

func TestMustGetDataFromJson(t *testing.T) {
	loader := NewDataTestLoader()

	rawBook := loader.getSingleRawBook

	if isIsbnNotEquals(rawBook.Isbn) {
		t.Fatal()
	}

	book := loader.getSingleBook

	if isIsbnNotEquals(book.Product_.Barcode) {
		t.Fatal()
	}

	rawBooks := loader.getRawBooks

	if isIsbnNotEquals(rawBooks[0].Isbn) {
		t.Fatal()
	}

	books := loader.getBooks

	if isIsbnNotEquals(books[0].Product_.Barcode) {
		t.Fatal()
	}

	dvd := loader.getSingleDvd

	if isUpacNotEquals(dvd.Product_.Barcode) {
		t.Fatal()
	}

	rawDvd := loader.getSingleRawDvd

	if isUpacNotEquals(rawDvd.Barcode) {
		t.Fatal()
	}

	dvds := loader.getDvds

	if isUpacNotEquals(dvds[0].Product_.Barcode) {
		t.Fatal()
	}

	rawDvds := loader.getRawDvds

	if isUpacNotEquals(rawDvds[0].Barcode) {
		t.Fatal()
	}
}
