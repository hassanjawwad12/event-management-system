package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {

	server.GET("/events", GetEvents)
	server.GET("/events/:id", GetEventById)

	// Authenticated Rputes
	server.POST("/events", CreateEvent)
	server.PUT("/events/:id", UpdateEvent)
	server.DELETE("/events/:id", DeleteEvent)

	server.POST("/signup", CreateUser)
	server.POST("/login", Login)
}
