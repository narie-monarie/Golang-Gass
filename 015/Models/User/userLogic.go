package models

import (
	"strconv"

	config "github.com/narie-monarie/Config"
	models "github.com/narie-monarie/Models"
)

type Person = models.Person

func GetPeople(count int) ([]Person, error) {
	rows, err := config.DB.Query(
		"SELECT id, first_name, last_name, email, password from people LIMIT " + strconv.Itoa(count),
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	people := make([]Person, 0)

	for rows.Next() {
		singlePerson := Person{}
		err = rows.Scan(
			&singlePerson.Id,
			&singlePerson.FirstName,
			&singlePerson.LastName,
			&singlePerson.Email,
			&singlePerson.Password,
		)
		if err != nil {
			return nil, err
		}
		people = append(people, singlePerson)
	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return people, err
}
