package main

import (
	"github.com/gin-gonic/gin"
	config "github.com/narie-monarie/Config"
	routes "github.com/narie-monarie/Routes"
)

func main() {
	//Fires up gin default router
	router := gin.Default()
	err := config.ConnectToDB()
	checkError(err)
	routes.UserRoutes(router)
	router.Run()
}
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
