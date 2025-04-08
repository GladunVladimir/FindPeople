package routes

import (
	"FindPeople/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/people", handlers.CreatePerson)
	r.GET("/people", handlers.GetAllPeople)
	r.GET("/people/:lastname", handlers.GetPersonByLastName)
	r.PUT("/people/:id", handlers.UpdatePerson)
}
