package main

import (
	"net/http"
	"strconv"

	"example.com/rest-api/db"
	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	server.GET("/event/:id", getEvent)
	server.Run(":8080")
}

// get an event by id
func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}
	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the event"})
		return
	}
	context.JSON(http.StatusOK, event)
}

// get all events available
func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong while getting the events"})
		return
	}
	context.JSON(http.StatusOK, events)
}

// crate an event
func createEvent(context *gin.Context) {
	var event models.Event
	// takes the data from post request and created a model out of it
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
		return
	}
	event.Id = 1
	event.CreatedBy = 1
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong while saving the events"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"Message": "Event created successfully", "event": event})
}
