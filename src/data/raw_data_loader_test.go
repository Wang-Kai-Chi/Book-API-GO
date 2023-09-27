package data

import (
	"testing"
)

func TestRawDataLoader(t *testing.T) {
	dt := NewDataForTest()
	rawBook := LoadData[RawBook]("../json/book_single.json")
	if rawBook.Isbn != dt.IsbnSample {
		t.Fatal()
	}
	t.Log(rawBook)

	rawBooks := LoadData[[]RawBook]("../json/book_arr.json")
	if rawBooks[0].Isbn != dt.IsbnSample {
		t.Fatal()
	}
	rawDvd := LoadData[RawDvd]("../json/dvd_single.json")
	if rawDvd.Barcode != dt.DvdCodeSample {
		t.Fatal()
	}

	rawDvds := LoadData[[]RawDvd]("../json/dvd_array.json")
	if rawDvds[0].Barcode != dt.DvdCodeSample {
		t.Fatal()
	}

	rawCd := LoadData[RawCd]("../json/cd_single.json")
	t.Log(rawCd)
	if rawCd.Barcode != dt.CdCodeSample {
		t.Fatal()
	}

	rawCds := LoadData[[]RawCd]("../json/cd_arr.json")
	if rawCds[0].Barcode != dt.CdCodeSample {
		t.Fatal()
	}

}
