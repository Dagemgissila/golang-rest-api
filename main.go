package main

import (
	"github.com/gin-gonic/gin"
	"restapi.com/dagem/db"
	"restapi.com/dagem/routes"
)

func main() {
	db.InitDB()
	server := gin.Default() //create default gin server
	routes.RegiterRoutes(server)
	server.Run(":8080")
}
