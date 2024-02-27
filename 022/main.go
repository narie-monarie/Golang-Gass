package main

import (
	"database/sql"
	"log"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

func main() {
	db, err := sql.Open("sqlite3", "./main.db")
	if err != nil {
		log.Fatal(err.Error())
	}

	boil.SetDB(db)
}
