package dvd

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	. "iknowbook.com/data"
	. "iknowbook.com/repository"
)

type DvdRepository struct {
	Connection *sqlx.DB
}

func NewDvdRepository(conn *sqlx.DB) DvdRepository {
	return DvdRepository{
		Connection: conn,
	}
}

func (serv DvdRepository) QueryWithLimit(limit int) []Dvd {
	return QueryEntity[Dvd](serv.Connection, NewDvdSqlStr().QueryWithLimit, limit)
}

func (serv DvdRepository) Insert(dvds []Dvd) sql.Result {
	return ExecSql[[]Dvd](serv.Connection, NewDvdSqlStr().Insert, dvds)
}
