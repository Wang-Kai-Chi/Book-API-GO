package handler

import (
	"testing"
)

const (
	UPAC = "4715219794386"
)

func GetSingleDvd() Dvd {
	return MustGetDvdFromJson(ReadFileAsString("../json/dvd_single.json"))
}

func GetDvds() []Dvd {
	return MustGetDvdArrayFromJson(ReadFileAsString("../json/dvd_array.json"))
}

func TestMustGetDvdFromJson(t *testing.T) {
	upac := GetSingleDvd().Barcode
	if upac != UPAC {
		t.Fatal()
	}
}

func TestMustGetDvdArrayFromJson(t *testing.T) {
	upac := GetDvds()[0].Barcode
	if upac != UPAC {
		t.Fatal()
	}
}
