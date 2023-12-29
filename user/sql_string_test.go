package user

import "testing"

func TestNewUserSqlStr(t *testing.T) {
	sql := NewUserSqlStr()

	t.Log(sql.QueryWithLimit)
	t.Log(sql.QueryByEmail)
	t.Log(sql.QueryByExactUserInfo)
	t.Log(sql.QueryById)
	t.Log(sql.UpdateUserAuth)
	t.Log(sql.Insert)
}
