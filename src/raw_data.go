package main

type RawData interface {
	RawBook | []RawBook | RawDvd | []RawDvd | RawCd | []RawCd
}

type RawBook struct {
	Isbn            string
	Title           string
	PublicationDate string
	Publisher       string
	Price           string
	Author          string
	Translator      string
	Language        string
}
type RawDvd struct {
	Barcode         string
	Title           string
	PublicationDate string
	Price           string
	Publisher       string
}
type RawCd struct {
	Barcode         string
	Title           string
	PublicationDate string
	Price           string
	Publisher       string
	Description     string
}
