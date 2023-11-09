package data

type Product struct {
	Product_id       int    `uri:"id" binding:"required"`
	Barcode          string `uri:"barcode" binding:"required"`
	Publication_date string `uri:"publication_date" binding:"required"`
	Product_title    string `uri:"title" binding:"required"`
	Price            string `uri:"price" binding:"required"`
	Publisher        string `uri:"publisher" binding:"required"`
	Quantity         int    `uri:"quantity" binding:"required"`
	Description      string `uri:"description" binding:"required"`
}
type Data interface {
	Product | []Product | Book | []Book | Dvd | []Dvd | Cd | []Cd
}
type Book struct {
	Product    `json:"Product"`
	Author     string
	Translator string
	Language   string
	Category   string
}

type Dvd struct {
	Product  `json:"Product"`
	Category string
	Director string
}
type Cd struct {
	Product   `json:"Product"`
	Performer string
	Genre     string
}
