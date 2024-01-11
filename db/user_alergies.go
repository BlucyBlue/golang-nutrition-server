package db

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

func AddUserAllergy(pool *pgxpool.Pool, userID int, allergyID int) error {
	ctx := context.Background()
	query := `INSERT INTO UserAllergies (UserID, AllergyID) VALUES ($1, $2)`
	_, err := pool.Exec(ctx, query, userID, allergyID)
	if err != nil {
		return err
	}
	return nil
}


func GetUserAllergies(pool *pgxpool.Pool, userID int) ([]int, error) {
	var allergies []int

	ctx := context.Background()

	query := `SELECT AllergyID FROM UserAllergies WHERE UserID = $1`
	rows, err := pool.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var allergyID int
		if err := rows.Scan(&allergyID); err != nil {
			return nil, err
		}
		allergies = append(allergies, allergyID)
	}

	return allergies, nil
}

func RemoveUserAllergy(pool *pgxpool.Pool, userID int, allergyID int) error {
	ctx := context.Background()

	query := `DELETE FROM UserAllergies WHERE UserID = $1 AND AllergyID = $2`
	_, err := pool.Exec(ctx, query, userID, allergyID)
	if err != nil {
		return err
	}
	return nil
}

