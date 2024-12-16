package models

import (
	"time"

	"example.com/rest-api/db"
)

type Event struct {
	Id          int64     `json:"id"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	StartTime   time.Time `json:"start_time" binding:"required"`
	CreatedBy   int       `json:"created_by"`
}

// a list of events
var events = []Event{}

// function to save an event and add it to the database
func (e Event) Save() error {
	// in the query, ? are used to denote that some values will be inserted here
	// and to prevent sql injection attacks
	query := `
	INSERT INTO events (title, description, location, start_time, created_by)
	VALUES (?,?,?,?,?)
	`

	// prepare the statement
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	// execute the statement
	result, err := stmt.Exec(e.Title, e.Description, e.Location, e.StartTime, e.CreatedBy)

	if err != nil {
		return err
	}

	// returns the last inserted primary key id
	id, err := result.LastInsertId()
	e.Id = id
	return err

}

// retrieve all events from the database
func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"

	/// DB.Query is uesd when fetching data from database
	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event
	// loop through all the rows to get populate the events list
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.Id, &event.Title, &event.Description, &event.Location, &event.StartTime, &event.CreatedBy)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

// fetch the event from the databse to fetch an event
func GetEventById(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.Id, &event.Title, &event.Description, &event.Location, &event.StartTime, &event.CreatedBy)
	if err != nil {
		return nil, err
	}
	return &event, nil

}

func (event Event) UpdateEvent() error {
	query := `
		UPDATE events 
		SET title = ?, description = ?, location = ?, start_time = ?, created_by = ?
		WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(event.Title, event.Description, event.Location, event.StartTime, event.CreatedBy, event.Id)
	return err
}
