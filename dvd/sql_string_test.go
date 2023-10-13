package dvd

import "testing"

func TestNewDvdSqlStr(t *testing.T) {
	sqlS := NewDvdSqlStr()
	t.Log(sqlS.QueryWithLimit)
	t.Log(sqlS.Insert)
}
