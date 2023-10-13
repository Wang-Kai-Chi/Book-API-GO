package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectDB() (*sqlx.DB, error) {
	config := MustGetConfig()
	db, err := sqlx.Connect(config.DriverName, config.DBConnection)

	return db, err
}
