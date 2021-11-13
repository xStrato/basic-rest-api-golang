package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xStrato/rest-api-30min/database"
	"github.com/xStrato/rest-api-30min/model"
	"github.com/xStrato/rest-api-30min/services"
)

func GetAllUsers(c *gin.Context) {
	db := database.GetConnection()
	var users []model.User
	if err := db.Find(&users).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Users cannot be listed: " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {
	db := database.GetConnection()
	user := model.NewUser()
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Cannot bind JSON: " + err.Error(),
		})
		return
	}
	services.Sha256Encrypt(&user.Password)
	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Cannot create user: " + err.Error(),
		})
		return
	}
	c.Status(http.StatusCreated)
}
