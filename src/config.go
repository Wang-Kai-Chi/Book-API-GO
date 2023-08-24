package main

import (
	"encoding/json"

	"iknowbook.com/handler"
)

type Config struct {
	DriverName   string
	DBConnection string
	SqlCFolder   string
}

func MustGetConfig() Config {
	var con Config
	raw := handler.MustReadFile("./resource/config.json")

	err := json.Unmarshal(raw, &con)
	if err != nil {
		panic(err)
	}
	return con
}
