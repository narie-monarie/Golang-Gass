package main

import (
	"narie/config"
	"narie/routes"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	err := config.Connect()
	if err != nil {
		panic(err)
	}
	routes.UserRoutes(router)
	http.ListenAndServe(":8090", router)

}
