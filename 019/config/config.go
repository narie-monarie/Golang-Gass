package config

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Connect() error {
	db, err := sql.Open("sqlite3", "./main.db")
	if err != nil {
		return err
	}
	// createTableSQL := `
	//        CREATE TABLE cats (
	//            id INTEGER PRIMARY KEY AUTOINCREMENT,
	//            catname TEXT,
	//            cattype TEXT
	//        );
	//    `
	// _, err = db.Exec(createTableSQL)
	// if err != nil {
	// 	panic(err)
	// }
	DB = db
	fmt.Println("Connected Succesfully")
	return nil
}
