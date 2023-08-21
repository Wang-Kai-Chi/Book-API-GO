package main

type Product struct {
	Id              int
	Barcode         string
	PublicationDate string
	Product_title   string
	Price           string
	Publisher       string
	Quantity        int
	Description     string
}

type Data interface {
	Book | []Book | Dvd | []Dvd | Cd | []Cd
}
type Book struct {
	Product_   Product
	Author     string
	Translator string
	Language   string
}

type Dvd struct {
	Product_ Product
	Category string
	Director string
}
type Cd struct {
	Product_  Product
	Performer string
	Genre     string
}