package data

import (
	"testing"
)

func TestRawDataConvertor(t *testing.T) {
	dt := NewDataForTest()

	var book Book
	book = book.ConvertRaw(LoadData[RawBook]("../json/book_single.json"))
	if book.Product.Barcode != dt.IsbnSample {
		t.Fatal()
	}
	t.Log(book)
	var books []Book
	books = book.ConvertRaws(LoadData[[]RawBook]("../json/book_arr.json"))
	if books[0].Product.Barcode != dt.IsbnSample {
		t.Fatal()
	}

	var dvd Dvd
	dvd = dvd.ConvertRaw(LoadData[RawDvd]("../json/dvd_single.json"))
	if dvd.Product.Barcode != dt.DvdCodeSample {
		t.Fatal()
	}
	var dvds []Dvd
	dvds = dvd.ConvertRaws(LoadData[[]RawDvd]("../json/dvd_array.json"))
	if dvds[0].Product.Barcode != dt.DvdCodeSample {
		t.Fatal()
	}

	var cd Cd
	cd = cd.ConvertRaw(LoadData[RawCd]("../json/cd_single.json"))
	if cd.Product.Barcode != dt.CdCodeSample {
		t.Fatal()
	}
	t.Log(cd)
	var cds []Cd
	cds = cd.ConvertRaws(LoadData[[]RawCd]("../json/cd_arr.json"))
	if cds[0].Product.Barcode != dt.CdCodeSample {
		t.Fatal()
	}
}
