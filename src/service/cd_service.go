package service

import (
	"github.com/jmoiron/sqlx"
	. "iknowbook.com/data"
)

type CdSqlStr struct {
	RelatedPath    string
	QueryWithLimit string
	Insert         string
}

func (sqlS *CdSqlStr) init() {
	sqlS.QueryWithLimit = GetSqlFromPath(sqlS.RelatedPath+sqlS.QueryWithLimit, sqlC)
	sqlS.Insert = GetSqlFromPath(sqlS.RelatedPath+sqlS.Insert, sqlC)
}

func NewCdSqlStr() *CdSqlStr {
	return NewSqlS[*CdSqlStr](`resource/sqlc/cd/cdSqlStr.json`, sqlC)
}

type CdService struct {
	Connection *sqlx.DB
}

func NewCdService(conn *sqlx.DB) CdService {
	return CdService{Connection: conn}
}

func (serv CdService) QueryWithLimit(limit int) []Cd {
	return QueryEntity[Cd](serv.Connection, NewCdSqlStr().QueryWithLimit, limit)
}

func (serv CdService) Insert(cds []Cd) {
	BulkExec[Cd](serv.Connection, NewCdSqlStr().Insert, cds)
}
