package service

import (
	"embed"
	"encoding/json"
)

type SqlS interface {
	*ProductSqlStr | *BookSqlStr | *DvdSqlStr | *CdSqlStr
	init()
}

//go:embed resource/sqlc/*
var sqlC embed.FS

func GetSqlFromPath(path string, fs embed.FS) string {
	data, _ := fs.ReadFile(path)
	return string(data)
}
func NewSqlS[T SqlS](mapPath string, fs embed.FS) T {
	data, err := fs.ReadFile(mapPath)
	var sqlS T
	if err == nil {
		err := json.Unmarshal(data, &sqlS)
		if err != nil {
			panic(err)
		}
		sqlS.init()
	} else {
		panic(err)
	}
	return sqlS
}
