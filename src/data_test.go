package main

import (
	"testing"
)

func isbnSample() string {
	return "9789571313887"
}

func dvdCodeSample() string {
	return "4715219794386"
}

func cdCodeSample() string {
	return "602508588662"
}

func isIsbnNotEquals(isbn string) bool {
	return isbn != isbnSample()
}

func isDvdCodeNotEquals(upac string) bool {
	return upac != dvdCodeSample()
}

func isCdCodeNotEqual(code string) bool {
	return code != cdCodeSample()
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
	if isDvdCodeNotEquals(dvd.Product_.Barcode) {
		t.Fatal()
	}

	rawDvd := loader.getSingleRawDvd
	if isDvdCodeNotEquals(rawDvd.Barcode) {
		t.Fatal()
	}

	dvds := loader.getDvds
	if isDvdCodeNotEquals(dvds[0].Product_.Barcode) {
		t.Fatal()
	}

	rawDvds := loader.getRawDvds
	if isDvdCodeNotEquals(rawDvds[0].Barcode) {
		t.Fatal()
	}

	cd := loader.getSingleCd
	if isCdCodeNotEqual(cd.Product_.Barcode) {
		t.Fatal()
	}

	cds := loader.getCds
	if isCdCodeNotEqual(cds[0].Product_.Barcode) {
		t.Fatal()
	}

	rawCd := loader.getSingleRawCd
	if isCdCodeNotEqual(rawCd.Barcode) {
		t.Fatal()
	}

	rawCds := loader.getRawCds
	if isCdCodeNotEqual(rawCds[0].Barcode) {
		t.Fatal()
	}
}
