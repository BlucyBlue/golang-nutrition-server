package main

import (
	"bytes"
	"context"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"github.com/jackc/pgx/v4/pgxpool"
)

var pool *pgxpool.Pool

func TestMain(m *testing.M) {
	var err error
	// Replace with your test database DSN
	dsn := "postgres://username:password@localhost:5432/testdb?sslmode=disable"
	pool, err = pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer pool.Close()

	os.Exit(m.Run())
}

func TestAddUserAllergyEndpoint(t *testing.T) {
	router := SetupRouter() // Ensure your router is using the pool

	// Prepare request body
	requestBody := bytes.NewBufferString(`{"userID": 1, "allergyID": 101}`)
	req, _ := http.NewRequest("POST", "/user/allergy/add", requestBody)
	req.Header.Set("Content-Type", "application/json")

	// Record the response
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	// Clean up test data
	pool.Exec(context.Background(), "DELETE FROM UserAllergies WHERE UserID = 1 AND AllergyID = 101")
}

func TestGetUserAllergiesEndpoint(t *testing.T) {
	router := SetupRouter() // Ensure your router is using the pool

	// Add test data (not shown for brevity)

	req, _ := http.NewRequest("GET", "/user/allergies/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	// Clean up test data
	pool.Exec(context.Background(), "DELETE FROM UserAllergies WHERE UserID = 1")
}


func TestDeleteUserAllergyEndpoint(t *testing.T) {
	router := SetupRouter() // Ensure your router is using the pool

	// Add test data (not shown for brevity)

	requestBody := bytes.NewBufferString(`{"userID": 1, "allergyID": 101}`)
	req, _ := http.NewRequest("POST", "/user/allergy/remove", requestBody)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

}

