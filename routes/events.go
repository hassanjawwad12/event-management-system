package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hassanjawwad12/event-management-system/models"
)

// GetEvents fetches all events from the database and returns them as a JSON response.
func GetEvents(context *gin.Context) {

	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch events"})
		return
	}
	context.JSON(http.StatusOK, events)
}

// CreateEvent creates a new event in the database and returns the created event as a JSON response.
func CreateEvent(context *gin.Context) {
	var event models.Event
	// store data from incoming body to event struct
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse Req Data"})
		return
	}
	event.ID = 1
	event.UserId = 1

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save event"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created successfully", "event": event})
}

// GetEventById fetches a specific event based on id
func GetEventById(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}
	event, err := models.GetEventByID(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch event"})
		return
	}
	context.JSON(http.StatusOK, event)
}
