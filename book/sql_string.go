package book

import (
	"embed"
)

type BookSqlStr struct {
	QueryByLimit      string
	QueryByConditions string
	Insert            string
	Update            string
}

//go:embed sql
var sqlC embed.FS

func MustReadFromPath(path string, fs embed.FS) string {
	data, _ := fs.ReadFile(path)
	return string(data)
}

func NewBookSqlStr() BookSqlStr {
	prefix := "sql/"

	return BookSqlStr{
		QueryByLimit:      MustReadFromPath(prefix+"query_book_by_limit.sql", sqlC),
		QueryByConditions: MustReadFromPath(prefix+"query_book_by_conditions.sql", sqlC),
		Insert:            MustReadFromPath(prefix+"insert_book.sql", sqlC),
		Update:            MustReadFromPath(prefix+"update_book.sql", sqlC),
	}
}
