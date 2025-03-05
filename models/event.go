package models

import (
	"fmt"
	"log"
	"time"

	"github.com/hassanjawwad12/event-management-system/db"
)

type Event struct {
	ID          int64
	Name        string    `json:"name"  binding:"required"`
	Description string    `json:"description"  binding:"required"`
	Location    string    `json:"location"  binding:"required"`
	DateTime    time.Time `json:"date_time"  binding:"required"`
	UserId      int64
}

// store it in the database
func (e *Event) Save() error {
	query := `
	INSERT INTO events (name, description, location, date_time, user_id) 
	VALUES (?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return fmt.Errorf("failed to prepare query: %v", err)
	}
	defer stmt.Close()
	// Exec is used to execute a prepared statement with the given arguments (updation)
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)
	if err != nil {
		return fmt.Errorf("failed to execute query: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get last insert ID: %v", err)
	}
	e.ID = id
	return err
}

// GetAllEvents returns all events from the database
func GetAllEvents() ([]Event, error) {
	// simple query thats why we did not prepare it
	query := `SELECT * FROM events`
	// Query is used to execute a query that returns rows, typically a SELECT.
	rows, err := db.DB.Query(query)
	fmt.Println("rows", rows)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch events: %v", err)
	}
	defer rows.Close()

	var events []Event

	// Next returns a boolean as long as there are rows in the result set.
	// It prepares the next result row for reading with the Scan method.
	for rows.Next() {
		var e Event
		err := rows.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserId)

		if err != nil {
			log.Printf("Skipping row due to error: %v", err)
			continue
		}

		events = append(events, e)
	}

	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("error iterating rows: %v", err)
	}

	return events, nil
}

// GetEventByID returns the event with the given ID
func GetEventByID(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

// Update updates the event in the database
func (event Event) Update() error {
	query := `
	UPDATE events
	SET name = ?, description = ?, location = ?, date_time = ?
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)
	return err
}

// Delete deletes the event with the given ID
func (event Event) Delete() error {
	query := "DELETE FROM events WHERE id = ?"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.ID)
	return err
}
