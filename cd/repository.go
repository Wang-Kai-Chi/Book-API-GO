package cd

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	. "iknowbook.com/data"
	. "iknowbook.com/repository"
)

type CdRepository struct {
	Connection *sqlx.DB
}

func NewCdRepository(conn *sqlx.DB) CdRepository {
	return CdRepository{Connection: conn}
}

func (serv CdRepository) QueryWithLimit(limit int) []Cd {
	return QueryEntity[Cd](serv.Connection, NewCdSqlStr().QueryWithLimit, limit)
}

func (serv CdRepository) Insert(cds []Cd) sql.Result {
	return ExecSql[[]Cd](serv.Connection, NewCdSqlStr().Insert, cds)
}
