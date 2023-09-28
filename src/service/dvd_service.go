package service

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
