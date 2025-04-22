package handlers

import (
	personapp "FindPeople/application/person"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type personReq struct {
	FullName string `json:"full_name"`
}
type friendReq struct {
	FriendID uint `json:"friend_id"`
}

func CreatePerson(c *gin.Context) {
	var req personReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	p, err := personapp.CreatePerson(personapp.CreatePersonInput{FullName: req.FullName})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, p)
}

func GetAllPeople(c *gin.Context) {
	if people, err := personapp.GetAllPeople(); err == nil {
		c.JSON(http.StatusOK, people)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

func GetPersonByLastName(c *gin.Context) {
	if res, err := personapp.FindByLastName(c.Param("lastname")); err == nil {
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

func UpdatePerson(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req personReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	if p, err := personapp.UpdatePerson(uint(id), req.FullName); err == nil {
		c.JSON(http.StatusOK, p)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

func AddFriend(c *gin.Context) {
	uid, _ := strconv.Atoi(c.Param("id"))
	var req friendReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	if err := personapp.LinkFriends(uint(uid), req.FriendID); err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "friends linked"})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

func GetFriends(c *gin.Context) {
	uid, _ := strconv.Atoi(c.Param("id"))
	if friends, err := personapp.GetFriends(uint(uid)); err == nil {
		c.JSON(http.StatusOK, friends)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
