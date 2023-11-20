package user

import (
	"embed"
	"log"
)

type UserSqlStr struct {
	QueryWithLimit  string
	QueryByUserInfo string

	Insert string
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

func NewUserSqlStr() UserSqlStr {
	prefix := "sql/"
	return UserSqlStr{
		QueryWithLimit:  mustReadFromPath(prefix+"query_with_limit.sql", sqlC),
		QueryByUserInfo: mustReadFromPath(prefix+"query_by_userinfo.sql", sqlC),

		Insert: mustReadFromPath(prefix+"insert.sql", sqlC),
	}
}
