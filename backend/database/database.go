package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

// InitDB initializes the database
func InitDB(filepath string) {
	var err error
	db, err = sql.Open("sqlite3", filepath)
	if err != nil {
		log.Fatal(err)
	}

	createTable := `
    CREATE TABLE IF NOT EXISTS urls (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        original_url TEXT NOT NULL,
        short_url TEXT NOT NULL UNIQUE
    );`

	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatal(err)
	}
}

// InsertURL inserts a new URL into the database
func InsertURL(originalURL, shortURL string) error {
	insertURL := `INSERT INTO urls (original_url, short_url) VALUES (?, ?);`
	_, err := db.Exec(insertURL, originalURL, shortURL)
	return err
}

// GetOriginalURL retrieves the original URL from the database using the short URL
func GetOriginalURL(shortURL string) (string, error) {
	var originalURL string
	query := `SELECT original_url FROM urls WHERE short_url = ?;`
	err := db.QueryRow(query, shortURL).Scan(&originalURL)
	return originalURL, err
}

func GetDB() *sql.DB {
	return db
}
