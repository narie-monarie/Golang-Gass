package routes

import (
	"narie/monarie/controllers"
	"net/http"
)

func CatRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /", controllers.GetCats)
}
