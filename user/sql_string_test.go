package user

import "testing"

func TestNewUserSqlStr(t *testing.T) {
	sql := NewUserSqlStr()

	t.Log(sql.QueryWithLimit)
	t.Log(sql.Insert)
}
