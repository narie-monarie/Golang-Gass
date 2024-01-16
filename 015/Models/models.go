package models

import (
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Person struct {
	Id        int        `json:"id"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Email     string     `json:"email"`
	IpAddress string     `json:"ip_address"`
	Product   *[]Product `json:"products"`
}

type Product struct {
	Id     int       `json:"id"`
	Name   string    `json:"product"`
	Price  float64   `json:"price"`
	Amount int64     `json:"amount_available"`
	Date   time.Time `json:"date_posted"`
	Status bool      `json:"is_available"`
}
