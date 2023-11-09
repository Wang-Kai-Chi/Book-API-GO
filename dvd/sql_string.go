package dvd

import "embed"

type DvdSqlStr struct {
	QueryWithLimit string
	Insert         string
}

//go:embed sql
var sqlC embed.FS

func MustReadFromPath(path string, fs embed.FS) string {
	data, _ := fs.ReadFile(path)
	return string(data)
}

func NewDvdSqlStr() DvdSqlStr {
	prefix := "sql/"

	return DvdSqlStr{
		QueryWithLimit: MustReadFromPath(prefix+"query_dvd_by_limit.sql", sqlC),
		Insert:         MustReadFromPath(prefix+"insert_dvd.sql", sqlC),
	}
}
