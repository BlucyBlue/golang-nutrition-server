package db

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

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
