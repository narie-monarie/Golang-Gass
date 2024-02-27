package config

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/narie-monarie/models"
	"github.com/sunary/sqlize"
)

var DB *sql.DB

func Connect() error {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		panic(err.Error())
	}
	DB = db
	return nil
}

func Migrate() error {
	newMigration := sqlize.NewSqlize(sqlize.WithSqlTag("sql"), sqlize.WithMigrationFolder("migrations"))
	_ = newMigration.FromObjects(models.User{})
	println(newMigration.StringUp())
	return nil
}
func AutoMigrate() error {
	newMigration := sqlize.NewSqlize(sqlize.WithSqlTag("sql"), sqlize.WithMigrationFolder("migrations"))
	err := newMigration.ArvoSchema(DB, models.User{})
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
