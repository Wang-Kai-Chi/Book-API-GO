package db

import (
	"embed"
	"encoding/json"
	"log"
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
			log.Println(err)
		}
	} else {
		log.Println(err)
	}
	return con
}
