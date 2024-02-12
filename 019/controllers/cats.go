package controllers

import (
	"database/sql"
	"encoding/json"
	"narie/config"
	"narie/models"
	"net/http"
)

type Cat = models.Cat

func getAllCats() ([]Cat, error) {
	rows, err := config.DB.Query("SELECT id,catname,cattype fROM cats")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var cats []Cat
	for rows.Next() {
		var cat Cat
		if err := rows.Scan(&cat.Id, &cat.Name, &cat.CatType); err != nil {
			return nil, err
		}
		cats = append(cats, cat)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return cats, nil
}

func GetCats(w http.ResponseWriter, r *http.Request) {
	cats, err := getAllCats()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cats)
}

// Get Cat By ID
func getCatByID(id string) (Cat, error) {
	stmt, err := config.DB.Prepare("SELECT id, catname, cattype FROM cats WHERE id = ?")
	if err != nil {
		return Cat{}, err
	}
	cat := Cat{}
	sqlErr := stmt.QueryRow(id).Scan(&cat.Id, &cat.Name, &cat.CatType)
	if sqlErr != nil {
		if sqlErr == sql.ErrNoRows {
			return Cat{}, sql.ErrNoRows
		}
		return Cat{}, sqlErr
	}
	return cat, nil
}

func GetCat(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	cat, err := getCatByID(id)

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Cat not found", http.StatusNoContent)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cat)
}

// Add Cat
