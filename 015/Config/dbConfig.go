package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

/*
const create string = `
CREATE TABLE Product (

	Id SERIAL PRIMARY KEY,
	Name VARCHAR(255),
	Price NUMERIC(10, 2),
	Amount BIGINT,
	Date TIMESTAMP,
	Status BOOLEAN,
	PersonId INT REFERENCES Person(Id)

);
`
*/
const (
	host     = "localhost"
	port     = 5432
	user     = "monari"
	password = "secret"
	dbname   = "narie"
)

func ConnectToDB() error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}

	/*
		if _, err := db.Exec(create); err != nil {
			return nil
		}
	*/

	DB = db
	fmt.Println("Successfully connected!")
	return nil
}
