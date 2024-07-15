package handlers

import (
	"log"
	"net/http"
	"url-shortener-backend/database"

	"github.com/gin-gonic/gin"
	"github.com/speps/go-hashids"
)

// ShortenURL handles the URL shortening request
func ShortenURL(c *gin.Context) {
	type Request struct {
		URL string `json:"url" binding:"required"`
	}
	var req Request
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	hd := hashids.NewData()
	h, _ := hashids.NewWithData(hd)
	shortURL, _ := h.Encode([]int{len(req.URL)})

	err := database.InsertURL(req.URL, shortURL)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to shorten URL"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"short_url": shortURL})
}

// RedirectURL handles the URL redirection request
func RedirectURL(c *gin.Context) {
	shortURL := c.Param("shortURL")

	originalURL, err := database.GetOriginalURL(shortURL)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	c.Redirect(http.StatusMovedPermanently, originalURL)
}

// package handlers

//import (
//	"net/http"

//"github.com/gin-gonic/gin"
//)

//func ShortenURL(c *gin.Context) {
//type Request struct {
//	URL string `json:"url"`
//}
//var req Request
//if err := c.BindJSON(&req); err != nil {
//	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
//	return
//}

// Logic to generate short URL
//shortURL := "shortened-url" // replace with actual logic

//c.JSON(http.StatusOK, gin.H{"short_url": shortURL})
//}

//func RedirectURL(c *gin.Context) {
// shortURL := c.Param("shortURL")

// Logic to retrieve original URL
//	originalURL := "http://example.com" // replace with actual logic

//	c.Redirect(http.StatusMovedPermanently, originalURL)
//}
