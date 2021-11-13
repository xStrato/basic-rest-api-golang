package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xStrato/rest-api-30min/database"
	"github.com/xStrato/rest-api-30min/model"
	"github.com/xStrato/rest-api-30min/services"
)

func Login(c *gin.Context) {
	db := database.GetConnection()
	login := model.Login{}
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Cannot bind JSON: " + err.Error(),
		})
		return
	}
	user := model.User{}
	if err := db.Where("email = ?", login.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Cannot find user: " + err.Error(),
		})
		return
	}
	services.Sha256Encrypt(&login.Password)
	if user.Password != login.Password {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid Credentiais",
		})
		return
	}
	token, err := services.NewJWTService().GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
