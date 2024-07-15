package main

import (
	"url-shortener-backend/database"
	"url-shortener-backend/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database
	database.InitDB("urls.db")

	// Set up the router
	router := gin.Default()

	// Define routes
	router.POST("/api/shorten", handlers.ShortenURL)
	router.GET("/:shortURL", handlers.RedirectURL)

	// Run the server
	router.Run(":8080")
}
