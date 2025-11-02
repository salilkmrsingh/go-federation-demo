package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func Connect() *sql.DB {
	db, err := sql.Open("sqlite3", "./user.db")
	if err != nil {
		log.Fatal(err)
	}

	// enable foreign keys
	_, err = db.Exec(`PRAGMA foreign_keys = ON;`)
	if err != nil {
		log.Fatal(err)
	}

	createUsers := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email Text NOT NULL
	);`

	if _, err := db.Exec(createUsers); err != nil {
		log.Fatalf("create users table: %v", err)
	}

	return db
}
