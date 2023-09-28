package service

import "testing"

func TestNewDvdSqlStr(t *testing.T) {
	sqlS := NewDvdSqlStr()
	t.Log(sqlS.RelatedPath)
	t.Log(sqlS.QueryWithLimit)
	t.Log(sqlS.Insert)
}
