package config

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func ConnectToDB() error {
	db, err := sql.Open("sqlite3", "../bass.sqlite")
	if err != nil {
		return err
	}
	DB = db
	return nil
}
