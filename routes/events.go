package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hassanjawwad12/event-management-system/models"
	"github.com/hassanjawwad12/event-management-system/utils"
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
	// check if the token is available in the header
	token := context.GetHeader("Authorization")
	if token == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Authorization token is required."})
		return
	}

	err := utils.VerifyToken(token)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

	var event models.Event
	err = context.ShouldBindJSON(&event) // store data from incoming body to event struct
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse Req Data"})
		return
	}
	event.ID = 1
	event.UserId = 1

	err = saveEvent(&event)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save event"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created successfully", "event": event})
}

func saveEvent(event *models.Event) error {
	return event.Save()
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

func UpdateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	_, err = models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the event."})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully!"})
}

func DeleteEvent(context *gin.Context) {

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	err = models.Delete(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete the event."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully!"})

}
