package main

import (
	"bytes"
	"context"
	"fmt"
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

func TestDeleteProductEndpoint(t *testing.T) {
	initTestDB()
	defer dbPool.Close()

	router := SetupRouter()

	productID, err := addDummyProductForTesting("Dummy Product", "Dummy Category")
	if err != nil {
		t.Fatalf("Error adding dummy product: %v", err)
	}

	req, _ := http.NewRequest("DELETE", fmt.Sprintf("/products/%d", productID), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

}

func TestGetProductByIDEndpoint(t *testing.T) {
	initTestDB()
	defer dbPool.Close()

	// Add a dummy product and get its ID
	productID, err := addDummyProductForTesting("Dummy Product", "Dummy Category")
	if err != nil {
		t.Fatalf("Error adding dummy product: %v", err)
	}

	// Create a request to get the product by ID
	req, _ := http.NewRequest("GET", fmt.Sprintf("/products/%d", productID), nil)
	w := httptest.NewRecorder()

	// Set up the router and serve the request
	router := SetupRouter()
	router.ServeHTTP(w, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Decode the response body into a product struct
	var product db.Product
	err = json.Unmarshal(w.Body.Bytes(), &product)
	if err != nil {
		t.Fatalf("Error decoding response: %v", err)
	}

	// Assert that the returned product matches the dummy product
	assert.Equal(t, productID, product.ProductID)
	assert.Equal(t, "Dummy Product", product.Name)
	assert.Equal(t, "Dummy Category", product.Category)

}
