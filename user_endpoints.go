package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"golang-nutrition-server/db"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
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

func Login(c *gin.Context) {
	var loginDetails struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&loginDetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
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

	// Retrieve the user from the database
	user, err := db.GetUserByUsername(pool, loginDetails.Username)
	if err != nil {
		// User not found or database error
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Compare the provided password with the stored hash
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginDetails.Password))
	if err != nil {
		// Passwords do not match
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	tokenValue := "generatedTokenValue" // This should be a securely generated token
	expiresIn := time.Now().Add(24 * time.Hour)

	err = db.CreateToken(pool, user.UserID, tokenValue, expiresIn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenValue})
}

func Logout(c *gin.Context) {
	tokenValue := c.GetHeader("Authorization") // Assuming the token is sent in the Authorization header

	// Connection string for your PostgreSQL database
	connString := "postgres://username:password@localhost:5432/dbname"

	// Establish a connection pool
	pool, err := pgxpool.Connect(context.Background(), connString)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer pool.Close()

	// Invalidate the token
	err = db.DeleteToken(pool, tokenValue)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error logging out"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}
