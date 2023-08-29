package main

type Product struct {
	Product_id       int
	Barcode          string
	Publication_date string
	Product_title    string `json:"title"`
	Price            string
	Publisher        string
	Quantity         int
	Description      string
}

type AData interface {
	Product | Book | Dvd | Cd
}
type Data interface {
	Product | []Product | Book | []Book | Dvd | []Dvd | Cd | []Cd
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
