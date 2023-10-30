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

func MustReadFromPath(path string, fs embed.FS) string {
	data, err := fs.ReadFile(path)
	if err != nil {
		log.Println(err)
	}
	return string(data)
}

func NewProductSqlStr() ProductSqlStr {
	prefix := "sql/"

	return ProductSqlStr{
		QueryWithLimit:      MustReadFromPath(prefix+"query_with_limit.sql", sqlC),
		QueryWithPriceRange: MustReadFromPath(prefix+"query_with_price_range.sql", sqlC),
		QueryByBarcode:      MustReadFromPath(prefix+"query_by_barcode.sql", sqlC),
		QueryByConditions:   MustReadFromPath(prefix+"query_by_conditions.sql", sqlC),
		QueryNewest:         MustReadFromPath(prefix+"query_newest.sql", sqlC),

		MaxPrice: MustReadFromPath(prefix+"max_price.sql", sqlC),

		Insert: MustReadFromPath(prefix+"insert.sql", sqlC),
		Update: MustReadFromPath(prefix+"update.sql", sqlC),
		Delete: MustReadFromPath(prefix+"delete.sql", sqlC),
	}
}
