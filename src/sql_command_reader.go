package main

import (
	"encoding/json"

	"iknowbook.com/handler"
)

type SqlC interface {
	ProductSqlC
	GetInsertString() string
}

func loadSqlC[T SqlC](path string) T {
	mustGetDataFromJson := func(content string) T {
		var entity T
		err := json.Unmarshal([]byte(content), &entity)
		if err != nil {
			panic(err)
		}
		return entity
	}
	return mustGetDataFromJson(handler.ReadFileAsString(path))
}

type ProductSqlC struct {
	InsertSQL string
}

func NewProductSqlC() ProductSqlC {
	config := MustGetConfig()

	return loadSqlC[ProductSqlC](config.SqlCFolder + "product_sqlc_map.json")
}

func (p ProductSqlC) GetInsertString() string {
	con := MustGetConfig()
	return handler.ReadFileAsString(string(con.SqlCFolder + p.InsertSQL))
}
