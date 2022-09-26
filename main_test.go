package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestGetAlbumsRoute(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/albums", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "[{\"id\":\"1\",\"title\":\"Blue Train\",\"artist\":\"John Coltrane\",\"price\":56.99},{\"id\":\"2\",\"title\":\"Jeru\",\"artist\":\"Gerry Mulligan\",\"price\":17.99},{\"id\":\"3\",\"title\":\"Sarah Vaughan and Clifford Brown\",\"artist\":\"Sarah Vaughan\",\"price\":39.99}]", w.Body.String())
}

func TestGetAlbumById(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/albums/2", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.JSONEq(t, 
		"{\"id\":\"2\",\"title\":\"Jeru\",\"artist\":\"Gerry Mulligan\",\"price\":17.99}", 
		w.Body.String())
}

func TestPostAlbum(t *testing.T) {
	router := setupRouter()
	var newAlbum = []byte(`{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}`)


	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/album", bytes.NewBuffer(newAlbum))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
	assert.Equal(t, 
		`{"id":"4","title":"The Modern Sound of Betty Carter","artist":"Betty Carter","price":49.99}`, 
		w.Body.String())
}