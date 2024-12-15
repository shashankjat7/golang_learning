package main

import (
	"net/http"

	"example.com/rest-api/db"
	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	server.Run(":8080")
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
