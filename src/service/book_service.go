package service

import (
	"embed"
	"encoding/json"
)

type BookSqlStr struct {
	RelatedPath  string
	Insert       string
	QueryByLimit string
}

//go:embed resource/sqlc/book/*
var bookSqlC embed.FS

func NewBookSqlStr() BookSqlStr {
	initBookSql := func(sqlS *BookSqlStr) {
		getSqlFromEmbededFolder := func(path string) string {
			data, _ := productSqlC.ReadFile(path)
			return string(data)
		}
		sqlS.QueryByLimit = getSqlFromEmbededFolder(sqlS.RelatedPath + sqlS.QueryByLimit)
		sqlS.Insert = getSqlFromEmbededFolder(sqlS.RelatedPath + sqlS.Insert)
	}
	data, err := bookSqlC.ReadFile("resource/sqlc/book/bookSqlStr.json")
	var sqlS BookSqlStr
	if err != nil {
		panic(err)
	} else {
		err := json.Unmarshal(data, &sqlS)
		if err != nil {
			panic(err)
		}
		initBookSql(&sqlS)
	}
	return sqlS
}
