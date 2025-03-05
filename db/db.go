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
}
