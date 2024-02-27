package routes

import (
	"narie/monarie/controllers"
	"net/http"
)

func CatRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /cats", controllers.GetCats)
	router.HandleFunc("GET /cat/{id}", controllers.GetCat)
	router.HandleFunc("POST /cat", controllers.GetCat)
}
