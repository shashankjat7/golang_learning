package models

import "time"

type Event struct {
	Id          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	StartTime   time.Time `json:"start_time"`
	CreatedBy   string    `json:"created_by"`
}

// a list of events
var events = []Event{}

// function to save an event and add it to the event list
func (e Event) save() {
	events = append(events, e)
}
