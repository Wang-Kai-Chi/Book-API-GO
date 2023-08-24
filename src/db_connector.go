package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectDB() (*sqlx.DB, error) {
	config := MustGetConfig()

	var dataBase *sqlx.DB
	var newErr error

	db, err := sqlx.Connect(config.DriverName, config.DBConnection)
	dataBase = db
	newErr = err

	return dataBase, newErr
}
