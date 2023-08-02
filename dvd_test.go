package main

import (
	"encoding/json"
	"testing"

	"iknowbook.com/handler"
)

func mustGetDvdFromJson(content string) Dvd {
	var dvd Dvd
	err := json.Unmarshal([]byte(content), &dvd)
	if err != nil {
		panic(err)
	}
	return dvd
}

func mustGetDvdArrayFromJson(content string) []Dvd {
	var dvds []Dvd
	err := json.Unmarshal([]byte(content), &dvds)
	if err != nil {
		panic(err)
	}
	return dvds
}

func upacSample() string {
	return "4715219794386"
}

func getSingleDvd() Dvd {
	return mustGetDvdFromJson(handler.ReadFileAsString("../json/dvd_single.json"))
}

func getDvds() []Dvd {
	return mustGetDvdArrayFromJson(handler.ReadFileAsString("../json/dvd_array.json"))
}

func TestMustGetDvdFromJson(t *testing.T) {
	upac := getSingleDvd().Barcode
	if upac != upacSample() {
		t.Fatal()
	}
}

func TestMustGetDvdArrayFromJson(t *testing.T) {
	upac := getDvds()[0].Barcode
	if upac != upacSample() {
		t.Fatal()
	}
}
