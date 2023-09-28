package data

type Product struct {
	Product_id       int
	Barcode          string
	Publication_date string
	Product_title    string
	Price            string
	Publisher        string
	Quantity         int
	Description      string
}
type Data interface {
	Product | []Product | Book | []Book | Dvd | []Dvd | Cd | []Cd
}
type Book struct {
	Product
	Author     string
	Translator string
	Language   string
	Category   string
}

type Dvd struct {
	Product
	Category string
	Director string
}
type Cd struct {
	Product
	Performer string
	Genre     string
}
