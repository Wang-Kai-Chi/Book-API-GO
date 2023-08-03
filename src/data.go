package main

type Book struct {
	Isbn            string
	Title           string
	PublicationDate string
	Price           string
	Author          string
	Translator      string
	Language        string
}

type Dvd struct {
	Barcode         string
	Title           string
	PublicationDate string
	Price           string
	Publisher       string
}

type Data interface {
	Book | []Book | Dvd | []Dvd
}
