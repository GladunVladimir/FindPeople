package routes

import (
	"FindPeople/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/people", handlers.CreatePerson)
	r.GET("/people", handlers.GetAllPeople)
	r.GET("/people/by-lastname/:lastname", handlers.GetPersonByLastName)
	r.PUT("/people/:id", handlers.UpdatePerson)
	r.POST("/people/:id/friends", handlers.AddFriend)
	r.GET("/people/:id/friends", handlers.GetFriends)
}
