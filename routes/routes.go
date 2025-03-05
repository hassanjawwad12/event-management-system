package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hassanjawwad12/event-management-system/middlewares"
)

func RegisterRoutes(server *gin.Engine) {

	server.GET("/events", GetEvents)
	server.GET("/events/:id", GetEventById)

	// Authenticated Rputes
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", CreateEvent)
	authenticated.PUT("/events/:id", UpdateEvent)
	authenticated.DELETE("/events/:id", DeleteEvent)

	server.POST("/signup", CreateUser)
	server.POST("/login", Login)
}
