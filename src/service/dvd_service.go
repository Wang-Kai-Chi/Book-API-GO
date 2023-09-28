package service

import (
	"github.com/jmoiron/sqlx"
	. "iknowbook.com/data"
)

type DvdSqlStr struct {
	RelatedPath    string
	QueryWithLimit string
	Insert         string
}

func (sqlS *DvdSqlStr) init() {
	sqlS.QueryWithLimit = GetSqlFromPath(sqlS.RelatedPath+sqlS.QueryWithLimit, sqlC)
	sqlS.Insert = GetSqlFromPath(sqlS.RelatedPath+sqlS.Insert, sqlC)
}

func NewDvdSqlStr() *DvdSqlStr {
	return NewSqlS[*DvdSqlStr](`resource/sqlc/dvd/dvdSqlStr.json`, sqlC)
}

type DvdService struct {
	Connection *sqlx.DB
}

func NewDvdService(conn *sqlx.DB) DvdService {
	return DvdService{
		Connection: conn,
	}
}

func (serv DvdService) QueryWithLimit(limit int) []Dvd {
	return QueryEntity[Dvd](serv.Connection, NewDvdSqlStr().QueryWithLimit, limit)
}

func (serv DvdService) Insert(dvds []Dvd) {
	BulkExec[Dvd](serv.Connection, NewDvdSqlStr().Insert, dvds)
}
