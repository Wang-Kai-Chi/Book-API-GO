package book

import "testing"

func TestNewBookSqlStr(t *testing.T) {
	sqlS := NewBookSqlStr()
	t.Log(sqlS.QueryByLimit)
	t.Log(sqlS.Insert)
	t.Log(sqlS.QueryByConditions)
	t.Log(sqlS.Update)
}
