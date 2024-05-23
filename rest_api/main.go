package main

import (
	"example.com/rest-api/controller"
	"example.com/rest-api/db"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitializeDB()
	server := gin.Default()

	controller.RegisterRoutes(server)

	server.Run(":8080") //localhost:8080 -> local domain.
}
