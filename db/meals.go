package db

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Meal struct {
	MealID      int
	Name        string
	Description string
}

// CreateMeal inserts a new meal into the database and returns its ID.
func CreateMeal(dbPool *pgxpool.Pool, name, description string) (int, error) {
	var mealID int
	err := dbPool.QueryRow(context.Background(), "INSERT INTO Meals (Name, Description) VALUES ($1, $2) RETURNING MealID", name, description).Scan(&mealID)
	if err != nil {
		return 0, err
	}
	return mealID, nil
}
