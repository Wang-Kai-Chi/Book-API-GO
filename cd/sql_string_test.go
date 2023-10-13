package cd

import "testing"

func TestNewCdSqlStr(t *testing.T) {
	sqlS := NewCdSqlStr()
	t.Log(sqlS.QueryWithLimit)
	t.Log(sqlS.Insert)
}
