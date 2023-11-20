package product

import (
	"embed"
	"log"
)

type ProductSqlStr struct {
	QueryWithLimit      string
	QueryWithPriceRange string
	QueryByBarcode      string
	QueryByConditions   string
	QueryNewest         string

	MaxPrice string

	Insert string
	Update string
	Delete string
}

//go:embed sql
var sqlC embed.FS

func mustReadFromPath(path string, fs embed.FS) string {
	data, err := fs.ReadFile(path)
	if err != nil {
		log.Println(err)
	}
	return string(data)
}

func NewProductSqlStr() ProductSqlStr {
	prefix := "sql/"

	return ProductSqlStr{
		QueryWithLimit:      mustReadFromPath(prefix+"query_with_limit.sql", sqlC),
		QueryWithPriceRange: mustReadFromPath(prefix+"query_with_price_range.sql", sqlC),
		QueryByBarcode:      mustReadFromPath(prefix+"query_by_barcode.sql", sqlC),
		QueryByConditions:   mustReadFromPath(prefix+"query_by_conditions.sql", sqlC),
		QueryNewest:         mustReadFromPath(prefix+"query_newest.sql", sqlC),

		MaxPrice: mustReadFromPath(prefix+"max_price.sql", sqlC),

		Insert: mustReadFromPath(prefix+"insert.sql", sqlC),
		Update: mustReadFromPath(prefix+"update.sql", sqlC),
		Delete: mustReadFromPath(prefix+"delete.sql", sqlC),
	}
}
