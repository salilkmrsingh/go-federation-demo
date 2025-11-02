package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func Connect() *sql.DB {
	db, err := sql.Open("sqlite3", "./review.db")
	if err != nil {
		log.Fatal(err)
	}

	// enable foreign keys
	_, err = db.Exec(`PRAGMA foreign_keys = ON;`)
	if err != nil {
		log.Fatal(err)
	}

	createReviews := `
	CREATE TABLE IF NOT EXISTS reviews (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
    	body TEXT,
    	author_id TEXT,
    	product_id TEXT
	);`

	if _, err := db.Exec(createReviews); err != nil {
		log.Fatalf("create users table: %v", err)
	}

	return db
}
