package service

import (
	"embed"
	"encoding/json"
)

type Config struct {
	DriverName   string
	DBConnection string
	SqlCFolder   string
}

//go:embed resource/config.json
var configFile embed.FS

func MustGetConfig() Config {
	raw, err := configFile.ReadFile("resource/config.json")

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
