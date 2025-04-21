package handlers

import (
	"FindPeople/database"
	"FindPeople/models"
	"FindPeople/service"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreatePerson создает нового пользователя по имени и связывает с внешними API.
// @Summary Создать человека
// @Tags People
// @Accept json
// @Produce json
// @Param person body models.CreatePersonInput true "Данные человека"
// @Success 201 {object} models.Person
// @Failure 400 {object} ErrorResponse
// @Router /people [post]
func CreatePerson(c *gin.Context) {
	var input models.CreatePersonInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid input"})
		return
	}
	firstName := strings.Split(input.FullName, " ")[0]
	age, _ := service.FetchAge(firstName)
	gender, _ := service.FetchGender(firstName)
	nationality, _ := service.FetchNationality(firstName)
	person := models.Person{
		FullName:    input.FullName,
		Gender:      gender,
		Nationality: nationality,
		Age:         age,
	}
	database.DB.Create(&person)
	c.JSON(http.StatusCreated, person)
}

// GetAllPeople возвращает список всех людей.
// @Summary Получить всех людей
// @Tags People
// @Produce json
// @Success 200 {array} models.Person
// @Router /people [get]
func GetAllPeople(c *gin.Context) {
	var people []models.Person
	database.DB.Find(&people)
	c.JSON(http.StatusOK, people)
}

// GetPersonByLastName ищет людей по фамилии.
// @Summary Найти человека по фамилии
// @Tags People
// @Produce json
// @Param lastname path string true "Фамилия"
// @Success 200 {array} models.Person
// @Router /people/by-lastname/{lastname} [get]
func GetPersonByLastName(c *gin.Context) {
	lastname := c.Param("lastname")
	var people []models.Person
	database.DB.Where("full_name ILIKE ?", "%"+lastname+"%").Find(&people)
	c.JSON(http.StatusOK, people)
}

// UpdatePerson обновляет данные человека.
// @Summary Обновить данные человека
// @Tags People
// @Accept json
// @Produce json
// @Param id path int true "ID человека"
// @Param person body models.Person true "Новые данные"
// @Success 200 {object} models.Person
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /people/{id} [put]
func UpdatePerson(c *gin.Context) {
	id := c.Param("id")
	var person models.Person
	if err := database.DB.First(&person, id).Error; err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Error: "Person not found"})
		return
	}
	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid data"})
		return
	}
	database.DB.Save(&person)
	c.JSON(http.StatusOK, person)
}

// AddFriend связывает двух людей как друзей.
// @Summary Добавить друга
// @Tags People
// @Accept json
// @Produce json
// @Param id path int true "ID пользователя"
// @Param friend body models.FriendInput true "ID друга"
// @Success 200 {object} MessageResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /people/{id}/friends [post]
func AddFriend(c *gin.Context) {
	var input models.FriendInput
	userID := c.Param("id")

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid input"})
		return
	}

	var person models.Person
	if err := database.DB.Preload("Friends").First(&person, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Error: "Person not found"})
		return
	}

	var friend models.Person
	if err := database.DB.First(&friend, input.FriendID).Error; err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Error: "Friend not found"})
		return
	}

	err := database.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&person).Association("Friends").Append(&friend); err != nil {
			return err
		}
		if err := tx.Model(&friend).Association("Friends").Append(&person); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, MessageResponse{Message: "Friends linked"})
}

// GetFriends возвращает список друзей пользователя.
// @Summary Получить друзей пользователя
// @Tags People
// @Produce json
// @Param id path int true "ID пользователя"
// @Success 200 {array} models.Person
// @Failure 404 {object} ErrorResponse
// @Router /people/{id}/friends [get]
func GetFriends(c *gin.Context) {
	userID := c.Param("id")
	var person models.Person

	if err := database.DB.Preload("Friends").First(&person, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Error: "Person not found"})
		return
	}

	c.JSON(http.StatusOK, person.Friends)
}
