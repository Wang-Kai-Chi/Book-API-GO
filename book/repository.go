package book

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	. "iknowbook.com/data"
	. "iknowbook.com/repository"
)

type BookRepository struct {
	Connection *sqlx.DB
}

func NewBookRepository(conn *sqlx.DB) BookRepository {
	return BookRepository{
		Connection: conn,
	}
}

func (serv BookRepository) mustGetRowsFromQuery(sqlStr string, params ...interface{}) *sqlx.Rows {
	return MustGetRowsFromQuery(serv.Connection, sqlStr, params...)
}

func (serv BookRepository) execSql(str string, es []Book) sql.Result {
	return ExecSql[[]Book](serv.Connection, str, es)
}

func (serv BookRepository) QueryBooks(db *sqlx.DB, sqlStr string, params ...interface{}) []Book {
	mustGetEntitiesFromRows := func(rows *sqlx.Rows) []Book {
		var entities []Book
		for rows.Next() {
			var p Product
			var b Book
			err := rows.Scan(&p.Product_id, &b.Author, &b.Translator,
				&b.Language, &b.Category, &p.Barcode, &p.Product_title,
				&p.Publisher, &p.Publication_date, &p.Quantity,
				&p.Description, &p.Price)
			if err != nil {
				panic(err)
			}
			b.Product = p
			entities = append(entities, b)
		}
		return entities
	}
	rows := serv.mustGetRowsFromQuery(sqlStr, params...)
	out := mustGetEntitiesFromRows(rows)
	return out
}

func (serv BookRepository) QueryWithLimit(limit int) []Book {
	return serv.QueryBooks(serv.Connection, NewBookSqlStr().QueryByLimit, limit)
}

func (serv BookRepository) Insert(books []Book) sql.Result {
	return serv.execSql(NewBookSqlStr().Insert, books)
}

func (serv BookRepository) QueryByConditions(pmin int, pmax int, book Book) []Book {
	return serv.QueryBooks(
		serv.Connection,
		NewBookSqlStr().QueryByConditions,
		pmin, pmax,
		book.Product_title, book.Publisher,
		book.Author, book.Translator,
		book.Language, book.Category,
	)
}

func (serv BookRepository) Update(books []Book) sql.Result {
	return serv.execSql(NewBookSqlStr().Update, books)
}
