// people_info_api.go

package main

import (
	"encoding/json"
	"gorm.io/driver/postgres"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Person struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	FullName    string `json:"full_name"`
	Gender      string `json:"gender"`
	Nationality string `json:"nationality"`
	Age         int    `json:"age"`
}

var db *gorm.DB

func createPerson(c *gin.Context) {
	var input struct {
		FullName string `json:"full_name"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	firstName := strings.Split(input.FullName, " ")[0]
	age, _ := fetchAge(firstName)
	gender, _ := fetchGender(firstName)
	nationality, _ := fetchNationality(firstName)

	person := Person{
		FullName:    input.FullName,
		Gender:      gender,
		Nationality: nationality,
		Age:         age,
	}
	db.Create(&person)
	c.JSON(http.StatusCreated, person)
}

func getAllPeople(c *gin.Context) {
	var people []Person
	db.Find(&people)
	c.JSON(http.StatusOK, people)
}

func getPersonByLastName(c *gin.Context) {
	lastname := c.Param("lastname")
	var people []Person
	db.Where("full_name ILIKE ?", "% "+lastname).Find(&people)
	c.JSON(http.StatusOK, people)
}

func updatePerson(c *gin.Context) {
	id := c.Param("id")
	var person Person
	if err := db.First(&person, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
		return
	}

	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}
	db.Save(&person)
	c.JSON(http.StatusOK, person)
}

func fetchAge(name string) (int, error) {
	resp, err := http.Get("https://api.agify.io?name=" + name)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var result struct {
		Age int `json:"age"`
	}
	json.NewDecoder(resp.Body).Decode(&result)
	return result.Age, nil
}

func fetchGender(name string) (string, error) {
	resp, err := http.Get("https://api.genderize.io?name=" + name)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result struct {
		Gender string `json:"gender"`
	}
	json.NewDecoder(resp.Body).Decode(&result)
	return result.Gender, nil
}

func fetchNationality(name string) (string, error) {
	resp, err := http.Get("https://api.nationalize.io?name=" + name)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result struct {
		Country []struct {
			CountryID string `json:"country_id"`
		} `json:"country"`
	}
	json.NewDecoder(resp.Body).Decode(&result)

	if len(result.Country) > 0 {
		return result.Country[0].CountryID, nil
	}
	return "", nil
}

func main() {
	var err error
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	db.AutoMigrate(&Person{})

	r := gin.Default()

	r.POST("/people", createPerson)
	r.GET("/people", getAllPeople)
	r.GET("/people/:lastname", getPersonByLastName)
	r.PUT("/people/:id", updatePerson)

	r.Run(":8080")
}
