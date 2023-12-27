package user

import (
	"embed"
	"log"
)

type UserSqlStr struct {
	QueryWithLimit       string
	QueryByUserInfo      string
	QueryByExactUserInfo string
	QueryById            string

	UpdateUserAuth string
	Insert         string
}

//go:embed sql
var sqlC embed.FS

func mustReadFromPath(path string, fs embed.FS) string {
	data, err := fs.ReadFile(path)
	if err != nil {
		log.Println(err)
	}
	return string(data)
}

func getSql(fileName string) string {
	prefix := "sql/"
	return mustReadFromPath(prefix+fileName, sqlC)
}

func NewUserSqlStr() UserSqlStr {
	return UserSqlStr{
		QueryWithLimit:       getSql("query_with_limit.sql"),
		QueryByUserInfo:      getSql("query_by_userinfo.sql"),
		QueryByExactUserInfo: getSql("query_by_exact_userinfo.sql"),
		QueryById:            getSql("query_by_id.sql"),

		UpdateUserAuth: getSql("update_user_auth.sql"),
		Insert:         getSql("insert.sql"),
	}
}
