package db

import (
	"database/sql"
	"log"
)

func ConnectUserSettings() *sql.DB {
	db, err := sql.Open("sqlite3", "./user_settings.db")
	if err != nil {
		log.Fatal(err)
	}

	// enable foreign keys
	_, err = db.Exec(`PRAGMA foreign_keys = ON;`)
	if err != nil {
		log.Fatal(err)
	}

	createUserSettings := `
	CREATE TABLE IF NOT EXISTS user_settings (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		theme TEXT DEFAULT 'light',
		language TEXT DEFAULT 'en',
		notifications_enabled BOOLEAN DEFAULT 1
	);`

	if _, err := db.Exec(createUserSettings); err != nil {
		log.Fatalf("create user_settings table: %v", err)
	}

	return db
}
