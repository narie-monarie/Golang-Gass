package config

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

/*

TODO
one to many in SQLITE

const create string = `

	  CREATE TABLE IF NOT EXISTS "people" (
		"id"	INTEGER,
		"first_name"	TEXT,
		"last_name"	TEXT,
		"email"	TEXT,
		"ip_address"	TEXT,
		PRIMARY KEY("id" AUTOINCREMENT)

);

		CREATE TABLE IF NOT EXISTS "people_products" (
	    "id" INTEGER PRIMARY KEY AUTOINCREMENT,
	    "person_id" INTEGER,
	    "product_id" INTEGER,
	    FOREIGN KEY ("person_id") REFERENCES "people" ("id"),
	    FOREIGN KEY ("product_id") REFERENCES "products" ("id")

);

`

const create string = `

	CREATE TABLE IF NOT EXISTS "products"(
		"id" INTEGER PRIMARY KEY AUTOINCREMENT,
		"name" TEXT,
		"price" INTEGER,
		"amount" INTEGER,
		"date" TEXT,
		"status" TEXT
	);

`
*/

const create string = `
	CREATE TABLE IF NOT EXISTS "people_products" (
	  "id" INTEGER PRIMARY KEY AUTOINCREMENT,
	  "person_id" INTEGER,
	  "product_id" INTEGER,
	  FOREIGN KEY ("person_id") REFERENCES "people" ("id"),
	  FOREIGN KEY ("product_id") REFERENCES "products" ("id")
	);
`

func ConnectToDB() error {
	db, err := sql.Open("sqlite3", "./tuts.db")
	if err != nil {
		return err
	}

	if _, err := db.Exec(create); err != nil {
		return nil
	}

	DB = db

	return nil
}
