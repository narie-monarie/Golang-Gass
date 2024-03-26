package main

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

// var schema = `
// DROP TABLE IF EXISTS user;
// CREATE TABLE user (
// 	user_id    INTEGER PRIMARY KEY,
//     first_name VARCHAR(80)  DEFAULT '',
//     last_name  VARCHAR(80)  DEFAULT '',
// 	email      VARCHAR(250) DEFAULT '',
// 	password   VARCHAR(250) DEFAULT NULL
// );
// `

type User struct {
	ID        int32 `sql:"primary_key;autoincrement"`
	FirstName string
	LastName  string
	Email     string
	Password  sql.NullString
}

var DB *sqlx.DB

func Connect() error {
	db, err := sqlx.Connect("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}
	DB = db
	return nil
}

func main() {
	app := fiber.New()
	if err := Connect(); err != nil {
		log.Fatal(err)
	}

	// newMigration := sqlize.NewSqlize(sqlize.WithSqlTag("sql"), sqlize.WithMigrationFolder("migrations"))
	// _ = newMigration.FromObjects(User{})
	// println(newMigration.StringUp())

	//db.MustExec(newMigration.StringUp())

	//tx := db.MustBegin()

	// tx.MustExec("INSERT INTO user (first_name, last_name, email, password) VALUES ($1, $2, $3,$4)", "Jason", "Moiron", "jmoiron@jmoiron.net", "kamatakamata")
	// tx.MustExec("INSERT INTO user (first_name, last_name, email, password) VALUES ($1, $2, $3,$4)", "John", "Doe", "johndoeDNE@gmail.net", "kamatakamata")
	// tx.NamedExec("INSERT INTO user (first_name, last_name, email, password) VALUES (:first_name, :last_name, :email, :password)",
	// 	&User{FirstName: "Jane", LastName: "Citizen", Email: "jane.citzen@example.com", Password: sql.NullString{String: "kamatakamata"}})
	// tx.Commit()
	// defer db.Close()

	app.Get("/users", GetAllUsers)
	app.Get("/user/:id", GetUser)
	app.Post("/user", CreateUser)
	app.Delete("/user/:id", DeleteUser)
	app.Patch("/user/:id", UpdateUser)
	app.Listen(":42069")
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user := User{}
	err := DB.Get(&user, "SELECT * FROM user WHERE user_id = ?", id)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(user)
}

func CreateUser(c *fiber.Ctx) error {
	user := User{}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Request",
		})
	}
	tx := DB.MustBegin()
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
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user := User{}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Request",
		})
	}
	tx := DB.MustBegin()
	tx.MustExec(
		"UPDATE user SET first_name = $1, last_name = $2, email = $3, password = $4 WHERE user_id = $?",
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password,
		id,
	)
	if err := tx.Commit(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to update user",
		})
	}
	return c.Status(fiber.StatusOK).JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	tx := DB.MustBegin()
	tx.MustExec("DELETE FROM user WHERE user_id = $?", id)
	if err := tx.Commit(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to delete user",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "user deleted",
	})
}

func GetAllUsers(c *fiber.Ctx) error {
	users := []User{}
	err := DB.Select(&users, "SELECT * FROM user")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to get users",
		})
	}
	return c.Status(fiber.StatusOK).JSON(users)
}
