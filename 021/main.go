package main

import (
	"narie/monarie/config"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	config.Connect()
	http.ListenAndServe(":8090", router)
}
