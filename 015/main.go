package main

import (
	"log"

	"github.com/gin-gonic/gin"
	config "github.com/narie-monarie/Config"
	routes "github.com/narie-monarie/Routes"
)

func main() {
	//Fires up gin default router
	router := gin.Default()
	err := config.ConnectToDB()
	if err != nil {
		log.Fatal(err)
	}
	routes.UserRoutes(router)
	router.Run()
}
