package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"golang-nutrition-server/db"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

type User struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func RegisterUser(c *gin.Context) {
	var newUser User
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}

	// Connection string for your PostgreSQL database
	connString := "postgres://username:password@localhost:5432/dbname"

	// Establish a connection pool
	pool, err := pgxpool.Connect(context.Background(), connString)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer pool.Close()

	err = db.SaveUserToDatabase(pool, newUser.Username, newUser.Email, string(hashedPassword))
	if err != nil {
		log.Printf("Error saving user: %v\n", err)
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}
