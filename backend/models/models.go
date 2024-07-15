package models

import (
	"database/sql"
	"errors"
	"math/rand"
	"net/url"
	"time"
	"url-shortener-backend/database"

	"github.com/mattn/go-sqlite3"
)

type URLInput struct {
	URL string `json:"url" binding:"required"`
}

func IsValidURL(inputURL string) bool {
	_, err := url.ParseRequestURI(inputURL)
	return err == nil
}

func GenerateShortURL() string {
	rand.Seed(time.Now().UnixNano())
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, 8)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func SaveURL(shortURL, originalURL string) error {
	db := database.GetDB()
	_, err := db.Exec("INSERT INTO urls (short_url, original_url) VALUES (?, ?)", shortURL, originalURL)
	if err != nil {
		if sqliteErr, ok := err.(sqlite3.Error); ok && sqliteErr.Code == sqlite3.ErrConstraint {
			return errors.New("URL already exists")
		}
		return err
	}
	return nil
}

func GetOriginalURL(shortURL string) (string, error) {
	db := database.GetDB()
	var originalURL string
	err := db.QueryRow("SELECT original_url FROM urls WHERE short_url = ?", shortURL).Scan(&originalURL)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("URL not found")
		}
		return "", err
	}
	return originalURL, nil
}
