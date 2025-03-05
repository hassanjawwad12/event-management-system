package main

import "github.com/gin-gonic/gin"

func main() {
	server := gin.Default()

	// takes a path and handler function
	server.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping pong",
		})
	})

	server.Run(":8080")
}
