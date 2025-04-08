package handlers

import (
	"FindPeople/database"
	"FindPeople/models"
	"FindPeople/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func CreatePerson(c *gin.Context) {
	var input struct {
		FullName string `json:"full_name"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	firstName := strings.Split(input.FullName, " ")[0]
	age, _ := services.FetchAge(firstName)
	gender, _ := services.FetchGender(firstName)
	nationality, _ := services.FetchNationality(firstName)

	person := models.Person{
		FullName:    input.FullName,
		Gender:      gender,
		Nationality: nationality,
		Age:         age,
	}
	database.DB.Create(&person)
	c.JSON(http.StatusCreated, person)
}

func GetAllPeople(c *gin.Context) {
	var people []models.Person
	database.DB.Find(&people)
	c.JSON(http.StatusOK, people)
}

func GetPersonByLastName(c *gin.Context) {
	lastname := c.Param("lastname")
	var people []models.Person
	database.DB.Where("full_name ILIKE ?", "% "+lastname).Find(&people)
	c.JSON(http.StatusOK, people)
}

func UpdatePerson(c *gin.Context) {
	id := c.Param("id")
	var person models.Person
	if err := database.DB.First(&person, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
		return
	}

	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}
	database.DB.Save(&person)
	c.JSON(http.StatusOK, person)
}
