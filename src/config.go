package main

import (
	"encoding/json"

	"iknowbook.com/handler"
)

type Config struct {
	DriverName   string
	DBConnection string
}

func GetConfig() (Config, error) {
	var con Config
	raw := handler.MustReadFile("../config.json")

	err := json.Unmarshal(raw, &con)
	return con, err
}
