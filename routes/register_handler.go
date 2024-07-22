package routes

import (
	"log"
	"net/http"
	"strconv"

	"dheek.com/restapi/models"
	"github.com/gin-gonic/gin"
)

func resgisterForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Printf("Error parsing id into int: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id."})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		log.Printf("Error getting event by id %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not get event by id"})
		return
	}

	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not register event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Registeration created."})

}

func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Printf("Error parsing id into int: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id."})
		return
	}

	var event models.Event
	event.ID = eventId

	err = event.CancelRegister(userId)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not delete event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Registeration cancelled successfully."})

}
