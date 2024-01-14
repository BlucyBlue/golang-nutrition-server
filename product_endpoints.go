package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"golang-nutrition-server/db"
	"log"
	"net/http"
	"strconv"
)

func AddProductEndpoint(c *gin.Context) {
	var product db.Product
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	connString := "postgres://username:password@localhost:5432/dbname"

	pool, err := pgxpool.Connect(context.Background(), connString)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer pool.Close()

	productID, err := db.AddProduct(pool, product.Name, product.Category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error adding product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product added successfully", "productID": productID})
}

func GetProductByIDEndpoint(c *gin.Context) {
	productID, err := strconv.Atoi(c.Param("productID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	connString := "postgres://username:password@localhost:5432/dbname"

	pool, err := pgxpool.Connect(context.Background(), connString)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer pool.Close()

	product, err := db.GetProductByID(pool, productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving product"})
		return
	}

	c.JSON(http.StatusOK, product)
}

func UpdateProductEndpoint(c *gin.Context) {
	var product db.Product
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	connString := "postgres://username:password@localhost:5432/dbname"

	pool, err := pgxpool.Connect(context.Background(), connString)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer pool.Close()

	err = db.UpdateProduct(pool, product.ProductID, product.Name, product.Category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}

func DeleteProductEndpoint(c *gin.Context) {
	productID, err := strconv.Atoi(c.Param("productID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}
	connString := "postgres://username:password@localhost:5432/dbname"
	// Establish a connection pool
	pool, err := pgxpool.Connect(context.Background(), connString)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer pool.Close()

	err = db.DeleteProduct(pool, productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}