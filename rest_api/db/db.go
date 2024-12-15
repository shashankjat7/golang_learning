package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// a global variable for databse
var DB *sql.DB

// initialize the database
func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect to database")
	}

	// how many connections should be open
	DB.SetMaxOpenConns(10)
	// how many max connections can be idle
	DB.SetMaxIdleConns(5)

	// created table
	createTables()
}

func createTables() {
	createEventsTableQuery := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		start_time DATETIME NOT NULL,
		created_by INTEGER NOTb NULL
	)
	`
	_, err := DB.Exec(createEventsTableQuery)
	if err != nil {
		panic(fmt.Sprint("An error occurred while creating the table", err))
	}
}
