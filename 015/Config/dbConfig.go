package config

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

/*
const create string = `
  CREATE TABLE IF NOT EXISTS "people" (
	"id"	INTEGER,
	"first_name"	TEXT,
	"last_name"	TEXT,
	"email"	TEXT,
	"ip_address"	TEXT,
	PRIMARY KEY("id" AUTOINCREMENT)
);
  `
*/

func ConnectToDB() error {
	db, err := sql.Open("sqlite3", "./tuts.db")
	if err != nil {
		return err
	}
	/*
		if _, err := db.Exec(create); err != nil {
			return nil
		}
	*/
	DB = db

	return nil
}
