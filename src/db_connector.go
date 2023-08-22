package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	config, err := GetConfig()

	var dataBase *sql.DB
	var newErr error

	if err == nil {
		db, err := sql.Open(config.DriverName, config.DBConnection)
		dataBase = db
		newErr = err
	}

	return dataBase, newErr
}
