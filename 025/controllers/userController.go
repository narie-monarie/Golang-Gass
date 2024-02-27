package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/narie-monarie/config"
	"github.com/narie-monarie/models"
)

type User = models.User

func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := User{}
	json.NewDecoder(r.Body).Decode(&user)
	newUser, err := config.DB.Query(
		"INSERT INTO users (username, password, email) VALUES (?, ?, ?)",
		user.Username,
		user.Password,
		user.Email,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if newUser.Err() != nil {
		http.Error(w, "cannot add user", http.StatusBadRequest)
		return
	}

	response := map[string]string{"message": "User added successfully"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	user := User{}
	json.NewDecoder(r.Body).Decode(&user)
}
