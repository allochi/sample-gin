package main

import (
	"github.com/allochi/sample-gin/handlers"
	"github.com/gin-gonic/gin"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetQuotes(t *testing.T) {
	router := gin.Default()
	router.GET("/quotes", handlers.GetQuotes)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/quotes", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	exp := "{\"name\":\"allochi\",\"age\":45}\n"
	assert.Equal(t, exp, w.Body.String())
}

func TestGetTenders(t *testing.T) {
	router := gin.Default()
	router.GET("/tenders", handlers.GetTenders)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/tenders", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	exp := "These is tenders!"
	assert.Equal(t, exp, w.Body.String())
}
