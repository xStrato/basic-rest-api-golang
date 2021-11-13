package server

import (
	"log"

	"github.com/xStrato/rest-api-30min/server/routes"

	"github.com/gin-gonic/gin"
)

type server struct {
	Port   string
	Server *gin.Engine
}

func New(port string) server {
	return server{
		Port:   port,
		Server: gin.Default(),
	}
}

func (s *server) Run() {
	router := routes.Configure(s.Server)

	log.Println("Server is running on port: " + s.Port)
	log.Fatalln(router.Run(":" + s.Port))
}
