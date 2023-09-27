package service

import (
	"database/sql"
	"embed"
	"encoding/json"

	"github.com/jmoiron/sqlx"
	. "iknowbook.com/data"
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
			data, _ := bookSqlC.ReadFile(path)
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

type BookService struct {
	Connection *sqlx.DB
}

func NewBookService(conn *sqlx.DB) BookService {
	return BookService{
		Connection: conn,
	}
}

func (serv BookService) QueryByLimit(limit int) []Book {
	return QueryEntity[Book](serv.Connection, NewBookSqlStr().QueryByLimit, limit)
}

func (serv BookService) Insert(books []Book) sql.Result {
	return ExecSql[Book](serv.Connection, NewBookSqlStr().Insert, books)
}
