package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectDB() (*sqlx.DB, error) {
	config, err := GetConfig()

	var dataBase *sqlx.DB
	var newErr error

	if err == nil {
		db, err := sqlx.Open(config.DriverName, config.DBConnection)
		dataBase = db
		newErr = err
	}

	return dataBase, newErr
}
