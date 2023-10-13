package cd

import "embed"

type CdSqlStr struct {
	QueryWithLimit string
	Insert         string
}

//go:embed sql
var sqlC embed.FS

func MustReadFromPath(path string, fs embed.FS) string {
	data, _ := fs.ReadFile(path)
	return string(data)
}

func NewCdSqlStr() CdSqlStr {
	prefix := "sql/"

	return CdSqlStr{
		QueryWithLimit: MustReadFromPath(prefix+"query_cd_by_limit.sql", sqlC),
		Insert:         MustReadFromPath(prefix+"insert_cd.sql", sqlC),
	}
}
