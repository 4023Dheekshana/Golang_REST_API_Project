package routes

import (
	"log"
	"net/http"
	"strconv"

	"dheek.com/restapi/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		log.Printf("Error fetching events: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events."})
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		log.Printf("Error binding JSON: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	userId := context.GetInt64("userId")
	event.UserID = userId

	err = event.Save()

	if err != nil {
		log.Printf("Error saving event: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event. Try again"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Printf("Error parsing id into int: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id."})
		return
	}
	eventid, err := models.GetEventById(eventId)
	if err != nil {
		log.Printf("Error fetching event by id: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not fetch event by id."})
		return
	}
	context.JSON(http.StatusOK, eventid)
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Printf("Error parsing id into int: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id."})
		return
	}
	userId := context.GetInt64("userId")
	event, err := models.GetEventById(eventId)
	if err != nil {
		log.Printf("Error getting event id: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not fetch event."})
		return
	}
	if event.UserID != userId {
		log.Fatalf("Error matching user id to allow access to update %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to authorize to update event"})
		return
	}
	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		log.Printf("Error binding JSON: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse updated data."})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.Update()
	if err != nil {
		log.Printf("Error updating event: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not update data."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Updated event successfully."})

}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Printf("Error parsing id into int: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id."})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		log.Printf("Error getting event id: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not fetch event."})
		return
	}
	err = event.Delete()
	if err != nil {
		log.Fatalf("Error deleting event: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not delete event."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully."})

}
