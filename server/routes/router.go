package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/xStrato/rest-api-30min/controllers"
	"github.com/xStrato/rest-api-30min/server/middlewares"
)

func Configure(gin *gin.Engine) *gin.Engine {
	main := gin.Group("api/v1")
	{
		login := main.Group("login")
		{
			login.POST("/", controllers.Login)
		}
		user := main.Group("user")
		{
			user.GET("/", controllers.GetAllUsers)
			user.POST("/", controllers.CreateUser)
		}
		books := main.Group("books", middlewares.Auth())
		{
			books.GET("/", controllers.GetAllBooks)
			books.GET("/:id", controllers.GetBookById)
			books.POST("/", controllers.CreateBook)
			books.PUT("/", controllers.UpdateBook)
			books.DELETE("/:id", controllers.DeleteBook)
		}
	}
	return gin
}
