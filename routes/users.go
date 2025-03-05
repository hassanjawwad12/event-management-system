package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hassanjawwad12/event-management-system/models"
)

func CreateUser(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse Req Data"})
		return
	}
	user.ID = 1
	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save user"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": user})
}
