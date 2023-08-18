package main

type Product struct {
	Id            int
	Barcode       string
	Product_title string
	Price         string
	Description   string
}

type Book struct {
	Product_        Product
	PublicationDate string
	Author          string
	Translator      string
	Language        string
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

type Dvd struct {
	Product_        Product
	PublicationDate string
	Publisher       string
}
type RawDvd struct {
	Barcode         string
	Title           string
	PublicationDate string
	Price           string
	Publisher       string
}

type Cd struct {
	Barcode         string
	Title           string
	PublicationDate string
	Price           string
	Publisher       string
	Description     string
}

type Data interface {
	Book | []Book | Dvd | []Dvd | RawDvd | []RawDvd | Cd | []Cd | RawBook | []RawBook
}
