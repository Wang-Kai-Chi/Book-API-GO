package db

import (
	"embed"
	"encoding/json"
)

type Config struct {
	DriverName   string
	DBConnection string
	SqlCFolder   string
}

//go:embed config.json
var configFile embed.FS

func MustGetConfig() Config {
	raw, err := configFile.ReadFile("config.json")

	var con Config
	if err == nil {
		err := json.Unmarshal(raw, &con)
		if err != nil {
			panic(err)
		}
	} else {
		panic(err)
	}
	return con
}
