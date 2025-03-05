package models

import "time"

type Event struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"  binding:"required"`
	Description string    `json:"description"  binding:"required"`
	Location    string    `json:"location"  binding:"required"`
	DateTime    time.Time `json:"date_time"  binding:"required"`
	UserId      int       `json:"user_id" `
}

var events = []Event{}

// store it in the database
func (e Event) Save() {
	events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}
