package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	router := setUpRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/healthcheck", nil)
	router.ServeHTTP(w, req)

	type aliveResponse struct {
		Alive bool `json:"alive"`
	}
	responseJSON, _ := json.Marshal(aliveResponse{Alive: true})

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, string(responseJSON), w.Body.String())
}
