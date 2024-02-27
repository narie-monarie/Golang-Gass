package controllers

import (
	"encoding/json"
	"narie/monarie/config"
	"narie/monarie/models"
	"net/http"
)

type Cat = models.Cat

func GetCats(w http.ResponseWriter, req *http.Request) {
	cats := Cat{}
	config.DB.Find(&cats)
	catjson, err := json.Marshal(&cats)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(catjson)
}
