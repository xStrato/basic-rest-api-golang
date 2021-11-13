package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/xStrato/rest-api-30min/controllers"
)

func Configure(gin *gin.Engine) *gin.Engine {
	main := gin.Group("api/v1")
	{
		books := main.Group("books")
		{
			books.GET("/", controllers.GetAll)
			books.GET("/:id", controllers.GetById)
			books.POST("/", controllers.Create)
			books.PUT("/", controllers.Update)
			books.DELETE("/:id", controllers.Delete)
		}
	}
	return gin
}
