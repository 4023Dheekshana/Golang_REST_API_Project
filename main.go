package main

import (
	"dheek.com/restapi/db"
	"dheek.com/restapi/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	defer db.DB.Close()
	server := gin.Default()

	routes.RegisterRoutes(server)
	server.Run(":8080") // localhost:8080
}
