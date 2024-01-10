package db

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

type User struct {
	UserID   int    `db:"UserID"`
	Username string `db:"Username"`
	Email    string `db:"Email"`
	Password string `db:"Password"`
}

func SaveUserToDatabase(pool *pgxpool.Pool, username, email, hashedPassword string) error {
	ctx := context.Background()

	// SQL statement to insert a new user
	query := `INSERT INTO Users (Username, Email, PasswordHash) VALUES ($1, $2, $3)`

	// Execute the query
	_, err := pool.Exec(ctx, query, username, email, hashedPassword)
	if err != nil {
		log.Printf("Error saving user to database: %v\n", err)
		return err
	}

	return nil
}

// GetUserByUsername retrieves a user by username from the database
func GetUserByUsername(pool *pgxpool.Pool, username string) (*User, error) {
	var user User
	err := pool.QueryRow(context.Background(), "SELECT UserID, Username, Email, PasswordHash FROM Users WHERE Username = $1", username).Scan(&user.Username, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
