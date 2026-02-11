package internal

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "./my.db")
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to SQLite")

	// Test connection
	var version string
	if err := db.QueryRow("SELECT sqlite_version()").Scan(&version); err != nil {
		return nil, err
	}

	fmt.Println("SQLite version:", version)

	// Create tables
	if err := CreateTables(db); err != nil {
		return nil, err
	}

	return db, nil
}

func CreateTables(db *sql.DB) error {
	createArticlesTable := `
	CREATE TABLE IF NOT EXISTS articles (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		content TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	if _, err := db.Exec(createArticlesTable); err != nil {
		return err
	}
	return nil
}
