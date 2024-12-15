package models

import "time"

type Event struct {
	Id          string    `json:"id"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	StartTime   time.Time `json:"start_time" binding:"required"`
	CreatedBy   string    `json:"created_by"`
}

// a list of events
var events = []Event{}

// function to save an event and add it to the event list
func (e Event) Save() {
	events = append(events, e)
}

// retrieve all events
func GetAllEvents() []Event {
	return events
}
