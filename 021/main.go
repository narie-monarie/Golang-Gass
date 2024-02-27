package main

import (
	"narie/monarie/config"
	"narie/monarie/routes"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	config.Connect()
	routes.CatRoutes(router)
	http.ListenAndServe(":8090", router)
}
