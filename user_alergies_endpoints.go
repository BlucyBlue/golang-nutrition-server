package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"net/http"
	"strconv"
	"golang-nutrition-server/db"
)


func AddUserAllergyEndpoint(c *gin.Context) {
	var input struct {
		UserID    int `json:"userID"`
		AllergyID int `json:"allergyID"`
	}
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	connString := "postgres://username:password@localhost:5432/dbname"

	// Establish a connection pool
	pool, err := pgxpool.Connect(context.Background(), connString)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer pool.Close()

	err = db.AddUserAllergy(pool, input.UserID, input.AllergyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error adding allergy to user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Allergy added to user successfully"})
}



func GetUserAllergiesEndpoint(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	connString := "postgres://username:password@localhost:5432/dbname"

	// Establish a connection pool
	pool, err := pgxpool.Connect(context.Background(), connString)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer pool.Close()

	allergies, err := db.GetUserAllergies(pool, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving user allergies"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"allergies": allergies})
}

func RemoveUserAllergyEndpoint(c *gin.Context) {
	var input struct {
		UserID    int `json:"userID"`
		AllergyID int `json:"allergyID"`
	}
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	connString := "postgres://username:password@localhost:5432/dbname"

	// Establish a connection pool
	pool, err := pgxpool.Connect(context.Background(), connString)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer pool.Close()

	err = db.RemoveUserAllergy(pool, input.UserID, input.AllergyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error removing allergy from user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Allergy removed from user successfully"})
}


