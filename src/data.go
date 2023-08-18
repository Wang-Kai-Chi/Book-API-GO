package main

type Product struct {
	Id            int
	Barcode       string
	Product_title string
	Price         string
	Description   string
}

type Data interface {
	Book | []Book | Dvd | []Dvd | Cd | []Cd
}
type Book struct {
	Product_        Product
	PublicationDate string
	Author          string
	Translator      string
	Language        string
}

type Dvd struct {
	Product_        Product
	PublicationDate string
	Publisher       string
}
type Cd struct {
	Product_        Product
	PublicationDate string
	Publisher       string
}

type RawData interface {
	RawBook | []RawBook | RawDvd | []RawDvd | RawCd | []RawCd
}

type RawBook struct {
	Isbn            string
	Title           string
	PublicationDate string
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
