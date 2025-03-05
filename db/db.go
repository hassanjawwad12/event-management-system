package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// InitDb initializes the database
var DB *sql.DB

func InitDb() {
	var err error
	DB, err = sql.Open("sqlite3", "event.db")
	if err != nil {
		panic("Could not connect to the database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	// create tables
	CreateTables()
}

func CreateTables() {

	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)
	`
	_, err := DB.Exec(createUsersTable)

	if err != nil {
		panic("Could not create users table.")
	}

	sqlStmt := `
	CREATE TABLE IF NOT EXISTS events (
	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, 
	name TEXT NOT NULL, 
	description TEXT NOT NULL, 
	location TEXT NOT NULL, 
	date_time DATETIME NOT NULL, 
	user_id INTEGER,
	FOREIGN KEY(user_id) REFERENCES users(id)
    )`

	// execute the SQL statement
	_, err = DB.Exec(sqlStmt)
	if err != nil {
		panic("Could not create table")
	}

	createRegistrationsTable := `
	CREATE TABLE IF NOT EXISTS registrations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		event_id INTEGER,
		user_id INTEGER,
		FOREIGN KEY(event_id) REFERENCES events(id),
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`
	_, err = DB.Exec(createRegistrationsTable)

	if err != nil {
		panic("Could not create registrations table.")
	}
}
