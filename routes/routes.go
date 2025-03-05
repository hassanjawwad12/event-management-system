package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {

	server.GET("/events", GetEvents)
	server.GET("/events/:id", GetEventById)
	server.POST("/events", CreateEvent)
}
