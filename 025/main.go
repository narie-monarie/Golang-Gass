package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sunary/sqlize"
)

var db *sql.DB

type User struct {
	Id       int64
	UserName string
	Password string
	Email    string
	Active   bool
	Phone    string
	Products *[]Product
}

type Product struct {
	Id          int64
	ProductName string
	Price       float64
	Available   bool
	Quantity    int
	UserId      int64
}

func main() {
	database, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		panic(err)
	}

	newMigration := sqlize.NewSqlize(sqlize.WithSqlTag("sql"), sqlize.WithMigrationFolder(""))
	_ = newMigration.FromObjects(User{})

	fmt.Println(newMigration.StringUp())

	fmt.Println("Database connected")
	db = database

}
