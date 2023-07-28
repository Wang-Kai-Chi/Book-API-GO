package handler

import "encoding/json"

type Dvd struct {
	Barcode         string
	Title           string
	PublicationDate string
	Price           string
	Publisher       string
}

func MustGetDvdFromJson(content string) Dvd {
	var dvd Dvd
	err := json.Unmarshal([]byte(content), &dvd)
	if err != nil {
		panic(err)
	}
	return dvd
}

func MustGetDvdArrayFromJson(content string) []Dvd {
	var dvds []Dvd
	err := json.Unmarshal([]byte(content), &dvds)
	if err != nil {
		panic(err)
	}
	return dvds
}
