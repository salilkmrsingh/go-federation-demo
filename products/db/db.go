package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func Connect() *sql.DB {
	db, err := sql.Open("sqlite3", "./product.db")
	if err != nil {
		log.Fatal(err)
	}

	// enable foreign keys
	_, err = db.Exec(`PRAGMA foreign_keys = ON;`)
	if err != nil {
		log.Fatal(err)
	}

	createProducts := `
	CREATE TABLE IF NOT EXISTS products (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		price REAL NOT NULL
	);`

	if _, err := db.Exec(createProducts); err != nil {
		log.Fatalf("create products table: %v", err)
	}

	return db
}
