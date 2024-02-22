package routes

import (
	"narie/controllers"
	"net/http"
)

func UserRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /", controllers.GetCats)
	router.HandleFunc("GET /{id}", controllers.GetCat)
	router.HandleFunc("POST /", controllers.AddCat)
}
