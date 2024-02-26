package main

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var schema = `
DROP TABLE IF EXISTS user;
CREATE TABLE user (
	user_id    INTEGER PRIMARY KEY,
    first_name VARCHAR(80)  DEFAULT '',
    last_name  VARCHAR(80)  DEFAULT '',
	email      VARCHAR(250) DEFAULT '',
	password   VARCHAR(250) DEFAULT NULL
);
`

type User struct {
	UserId    int    `db:"user_id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string
	Password  sql.NullString
}

func main() {
	app := fiber.New()
	db, err := sqlx.Connect("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}
	db.MustExec(schema)

	tx := db.MustBegin()

	tx.MustExec("INSERT INTO user (first_name, last_name, email, password) VALUES ($1, $2, $3,$4)", "Jason", "Moiron", "jmoiron@jmoiron.net", "kamatakamata")
	tx.MustExec("INSERT INTO user (first_name, last_name, email, password) VALUES ($1, $2, $3,$4)", "John", "Doe", "johndoeDNE@gmail.net", "kamatakamata")
	tx.NamedExec("INSERT INTO user (first_name, last_name, email, password) VALUES (:first_name, :last_name, :email, :password)",
		&User{FirstName: "Jane", LastName: "Citizen", Email: "jane.citzen@example.com", Password: sql.NullString{String: "kamatakamata"}})
	tx.Commit()
	defer db.Close()

	app.Get("/users", func(c *fiber.Ctx) error {
		users := []User{}
		err := db.Select(&users, "SELECT * FROM user")
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "failed to get users",
			})
		}
		return c.Status(fiber.StatusOK).JSON(users)
	})

	app.Get("/user/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		user := User{}
		err := db.Get(&user, "SELECT * FROM user WHERE user_id = ?", id)

		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "User not found",
			})
		}
		return c.Status(fiber.StatusOK).JSON(user)
	})

	app.Post("/user", func(c *fiber.Ctx) error {
		user := User{}
		if err := c.BodyParser(&user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid Request",
			})
		}
		tx := db.MustBegin()
		tx.MustExec(
			"INSERT INTO user (first_name, last_name, email, password) VALUES ($1, $2, $3,$4)",
			user.FirstName,
			user.LastName,
			user.Email,
			user.Password,
		)
		if err := tx.Commit(); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "failed to create new user",
			})
		}
		return c.Status(fiber.StatusCreated).JSON(user)
	})
	app.Listen(":3000")
}
