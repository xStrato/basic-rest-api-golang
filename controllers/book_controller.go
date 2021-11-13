package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/xStrato/rest-api-30min/database"
	"github.com/xStrato/rest-api-30min/model"
)

func GetAll(c *gin.Context) {
	db := database.GetConnection()
	var books []model.Book
	if err := db.Find(&books).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Books cannot be listed: " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, books)
}

func GetById(c *gin.Context) {
	if id := c.Param("id"); strings.TrimSpace(id) == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "'id' cannot be null or empty",
		})
		return
	}
	db := database.GetConnection()
	var book model.Book
	if err := db.First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Book cannot be found: " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, book)
}

func Create(c *gin.Context) {
	db := database.GetConnection()
	var book model.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Cannot bind JSON: " + err.Error(),
		})
		return
	}
	if err := db.Create(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Cannot create book: " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, book)
}
