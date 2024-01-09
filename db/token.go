package db

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"time"
)

// Token represents a user authentication token
type Token struct {
	TokenID   int
	UserID    int
	Token     string
	CreatedAt time.Time
	ExpiresAt time.Time
}

// CreateToken creates a new token for a given user
func CreateToken(pool *pgxpool.Pool, userID int, token string, expiresAt time.Time) error {
	ctx := context.Background()
	query := `INSERT INTO Tokens (UserID, Token, ExpiresAt) VALUES ($1, $2, $3)`
	_, err := pool.Exec(ctx, query, userID, token, expiresAt)
	if err != nil {
		log.Printf("Error creating token: %v\n", err)
		return err
	}
	return nil
}

// FindToken searches for a token by its value
func FindToken(pool *pgxpool.Pool, tokenValue string) (*Token, error) {
	ctx := context.Background()
	token := &Token{}
	query := `SELECT TokenID, UserID, Token, CreatedAt, ExpiresAt FROM Tokens WHERE Token = $1`
	err := pool.QueryRow(ctx, query, tokenValue).Scan(&token.TokenID, &token.UserID, &token.Token, &token.CreatedAt, &token.ExpiresAt)
	if err != nil {
		log.Printf("Error finding token: %v\n", err)
		return nil, err
	}
	return token, nil
}

// DeleteToken deletes a token
func DeleteToken(pool *pgxpool.Pool, tokenValue string) error {
	ctx := context.Background()
	query := `DELETE FROM Tokens WHERE Token = $1`
	_, err := pool.Exec(ctx, query, tokenValue)
	if err != nil {
		log.Printf("Error deleting token: %v\n", err)
		return err
	}
	return nil
}
