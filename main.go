package main

import (
	"github.com/xStrato/rest-api-30min/database"
	"github.com/xStrato/rest-api-30min/server"
)

func main() {
	database.StartDB()
	server := server.New("3000")
	server.Run()
}
