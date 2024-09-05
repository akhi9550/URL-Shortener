package main_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	main "urlshortener"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/shorten", main.ShortenURLHandler)
	r.GET("/:shortUrl", main.RetrieveURLHandler)
	return r
}

func TestShortenURLHandler(t *testing.T) {
	router := setupRouter()
	body := `{"url": "http://example.com"}`
	req, _ := http.NewRequest(http.MethodPost, "/shorten", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "short_url")
	assert.Contains(t, w.Body.String(), main.BaseURL)
}

func TestRetrieveURLHandler(t *testing.T) {
	router := setupRouter()
	testShortURL := "abcde"
	originalURL := "http://example.com"
	main.Mu.Lock()
	main.UrlStore[testShortURL] = originalURL
	main.Mu.Unlock()

	req, _ := http.NewRequest(http.MethodGet, "/"+testShortURL, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusPermanentRedirect, w.Code)
	assert.Equal(t, originalURL, w.Header().Get("Location"))

	req, _ = http.NewRequest(http.MethodGet, "/invalid", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), "URL not found")
}
