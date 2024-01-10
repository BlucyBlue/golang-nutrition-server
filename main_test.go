package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRootEndpoint(t *testing.T) {
	router := SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"message":"Server is up!"}`, w.Body.String())
}

func TestRegisterUser(t *testing.T) {
	// Get the router with the registered routes
	router := SetupRouter()

	// Create a request body with user data
	requestBody := bytes.NewBufferString(`{"username": "testuser", "email": "test@example.com", "password": "testpass"}`)

	// Create a request to send to the above route
	req, err := http.NewRequest(http.MethodPost, "/register", requestBody)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Record the response
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "User registered successfully")
}

func TestLogin(t *testing.T) {
	// Set up the router
	router := SetupRouter()

	// Create a request body
	requestBody := bytes.NewBufferString(`{"username": "testuser", "password": "testpass"}`)

	// Create a request
	req, _ := http.NewRequest("POST", "/login", requestBody)
	req.Header.Set("Content-Type", "application/json")

	// Record the response
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Check the response
	assert.Equal(t, http.StatusOK, w.Code)
	// Further checks can be added to validate the response content, such as the presence of a token
}

func TestLogout(t *testing.T) {
	// Set up the router
	router := SetupRouter()

	// Create a request
	req, _ := http.NewRequest("POST", "/logout", nil)
	// Set a valid token in the Authorization header
	req.Header.Set("Authorization", "validToken")

	// Record the response
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Check the response
	assert.Equal(t, http.StatusOK, w.Code)
	// Further checks can be added based on the logout logic
}
