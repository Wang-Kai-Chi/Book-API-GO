package repository

import (
	"database/sql"
	"log"

	"github.com/jmoiron/sqlx"
	. "iknowbook.com/app/data"
)

func MustGetRowsFromQuery(db *sqlx.DB, sqlStr string, params ...interface{}) *sqlx.Rows {
	query, args, err := sqlx.In(sqlStr, params...)
	if err != nil {
		panic(err)
	}
	query = db.Rebind(query)
	rows, err := db.Queryx(query, args...)
	if err != nil {
		panic(err)
	}
	return rows
}

func QueryEntity[T Data](db *sqlx.DB, sqlStr string, params ...interface{}) []T {
	mustGetEntitiesFromRows := func(rows *sqlx.Rows) []T {
		var entities []T
		for rows.Next() {
			var p T
			err := rows.StructScan(&p)
			if err != nil {
				panic(err)
			}
			entities = append(entities, p)
		}
		return entities
	}
	rows := MustGetRowsFromQuery(db, sqlStr, params...)
	out := mustGetEntitiesFromRows(rows)
	return out
}

func ExecSql[T Data](db *sqlx.DB, str string, ps T) sql.Result {
	res, err := db.NamedExec(str, ps)
	if err != nil {
		panic(err)
	}
	return res
}

func BulkExec[T Data](db *sqlx.DB, str string, es []T) {
	for _, v := range es {
		_, err := db.NamedExec(str, v)
		if err != nil {
			log.Println(err)
		}
	}
}
