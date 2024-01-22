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

// GetMealByID fetches a meal by its ID from the database
func GetMealByID(dbPool *pgxpool.Pool, mealID int) (*Meal, error) {
	meal := &Meal{}

	err := dbPool.QueryRow(context.Background(), "SELECT MealID, Name, Description FROM Meals WHERE MealID = $1", mealID).Scan(&meal.MealID, &meal.Name, &meal.Description)
	if err != nil {
		return nil, err
	}

	return meal, nil
}

// GetAllMeals fetches all meals from the database
func GetAllMeals(dbPool *pgxpool.Pool) ([]Meal, error) {
	var meals []Meal

	rows, err := dbPool.Query(context.Background(), "SELECT MealID, Name, Description FROM Meals")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var meal Meal
		if err := rows.Scan(&meal.MealID, &meal.Name, &meal.Description); err != nil {
			return nil, err
		}
		meals = append(meals, meal)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return meals, nil
}
