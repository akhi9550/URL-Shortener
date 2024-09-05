package main

import (
	"math/rand"
	"net/http"
	"net/url"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	UrlStore = make(map[string]string)
	Mu       sync.RWMutex
	BaseURL  = "http://localhost:8080/"
)

func main() {
	r := gin.Default()
	r.POST("/shorten", ShortenURLHandler)
	r.GET("/:shortUrl", RetrieveURLHandler)
	r.Run(":8080")
}

type requestBody struct {
	URL string `json:"url" binding:"required"`
}

func ShortenURLHandler(c *gin.Context) {
	var request requestBody
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request"})
		return
	}
	if !isValidURL(request.URL) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL"})
		return
	}

	shortURL := generateShortURL()

	Mu.Lock()
	UrlStore[shortURL] = request.URL
	Mu.Unlock()

	c.JSON(http.StatusOK, gin.H{"short_url": BaseURL + shortURL})
}

func RetrieveURLHandler(c *gin.Context) {
	shortURL := c.Param("shortUrl")
	Mu.RLock()
	originalURL, exists := UrlStore[shortURL]
	Mu.RUnlock()

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	c.Redirect(http.StatusPermanentRedirect, originalURL)
}

func isValidURL(testURL string) bool {
	_, err := url.ParseRequestURI(testURL)
	return err == nil
}

func generateShortURL() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const shortURLLength = 5
	shortURL := make([]byte, shortURLLength)
	for i := range shortURL {
		shortURL[i] = charset[rand.Intn(len(charset))]
	}
	return string(shortURL)
}
