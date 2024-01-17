package main

import (
	"bytes"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"golang-nutrition-server/db"
	"encoding/json"
	"github.com/stretchr/testify/assert"
)

var dbPool *pgxpool.Pool

func initTestDB() {
	var err error
	dsn := "postgres://username:password@localhost:5432/testdb?sslmode=disable"
	dbPool, err = pgxpool.Connect(context.Background(),

		dsn)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}


}

func tearDownTestDB() {
	dbPool.Close()
}

func insertDummyProduct(name, category string) (int, error) {
	var productID int
	err := dbPool.QueryRow(context.Background(), "INSERT INTO Products (Name, Category) VALUES ($1, $2) RETURNING ProductID", name, category).Scan(&productID)
	return productID, err
}

func deleteDummyProduct(productID int) {
	dbPool.Exec(context.Background(), "DELETE FROM Products WHERE ProductID = $1", productID)
}

func addDummyProductForTesting(name, category string) (int, error) {
	var productID int
	err := dbPool.QueryRow(context.Background(), "INSERT INTO Products (Name, Category) VALUES ($1, $2) RETURNING ProductID", name, category).Scan(&productID)
	if err != nil {
		return 0, err
	}
	return productID, nil
}

func TestAddProductEndpoint(t *testing.T) {
	initTestDB()
	defer tearDownTestDB()

	router := SetupRouter()

	product := db.Product{Name: "Test Product", Category: "Test Category"}
	productJSON, _ := json.Marshal(product)
	requestBody := bytes.NewBuffer(productJSON)

	req, _ := http.NewRequest("POST", "/products", requestBody)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

}

func TestUpdateProductEndpoint(t *testing.T) {
	initTestDB()
	defer dbPool.Close()

	productID, err := addDummyProductForTesting("Dummy Product", "Dummy Category")
	if err != nil {
		t.Fatalf("Error adding dummy product: %v", err)
	}

	router := SetupRouter()

	updatedProduct := db.Product{ProductID: productID, Name: "Updated Name", Category: "Updated Category"} // Use actual product ID
	productJSON, _ := json.Marshal(updatedProduct)
	requestBody := bytes.NewBuffer(productJSON)

	req, _ := http.NewRequest("PUT", "/products", requestBody)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

}