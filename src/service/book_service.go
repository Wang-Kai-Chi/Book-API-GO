package service

import (
	"github.com/jmoiron/sqlx"
	. "iknowbook.com/data"
)

type BookSqlStr struct {
	RelatedPath  string
	Insert       string
	QueryByLimit string
}

func (sqlS *BookSqlStr) init() {
	sqlS.QueryByLimit = GetSqlFromPath(sqlS.RelatedPath+sqlS.QueryByLimit, sqlC)
	sqlS.Insert = GetSqlFromPath(sqlS.RelatedPath+sqlS.Insert, sqlC)
}

func NewBookSqlStr() *BookSqlStr {
	return NewSqlS[*BookSqlStr](`resource/sqlc/book/bookSqlStr.json`, sqlC)
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

func (serv BookService) Insert(books []Book) {
	BulkExec[Book](serv.Connection, NewBookSqlStr().Insert, books)
}
