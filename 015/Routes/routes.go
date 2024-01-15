package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/narie-monarie/Controllers"
)

func UserRoutes(router *gin.Engine) {
	user := router.Group("/api/v1")
	{
		user.GET("person", controllers.GetPersons)
		user.GET("person/:id", controllers.GetPersonById)
		user.POST("person", controllers.AddPerson)
		user.PUT("person/:id", controllers.UpdatePerson)
		user.DELETE("person/:id", controllers.DeletePerson)
		user.OPTIONS("person", controllers.Options)
	}
}
