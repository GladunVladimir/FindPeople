package main

import (
	"FindPeople/database"
	"FindPeople/models"
	"FindPeople/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	database.ConnectDB()
	err := database.DB.AutoMigrate(&models.Person{})
	if err != nil {
		return
	}
	routes.SetupRoutes(r)
	err = r.Run(":8080")
	if err != nil {
		return
	}
}
