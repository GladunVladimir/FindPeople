package main

import (
	"FindPeople/database"
	_ "FindPeople/docs"
	"FindPeople/models"
	"FindPeople/routes"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	r := gin.Default()
	database.ConnectDB()
	err := database.DB.AutoMigrate(&models.Person{})
	if err != nil {
		return
	}
	routes.SetupRoutes(r)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	err = r.Run(":8080")
	if err != nil {
		return
	}
}
