package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hassanjawwad12/event-management-system/db"
	"github.com/hassanjawwad12/event-management-system/routes"
)

func main() {
	db.InitDb()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
