package db

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

func AddAllergyToProduct(dbPool *pgxpool.Pool, productID, allergyID int) error {
	_, err := dbPool.Exec(context.Background(), "INSERT INTO AllergyProducts (ProductID, AllergyID) VALUES ($1, $2)", productID, allergyID)
	if err != nil {
		return err
	}
	return nil
}

func GetAllergiesForProduct(dbPool *pgxpool.Pool, productID int) ([]int, error) {
	var allergyIDs []int

	rows, err := dbPool.Query(context.Background(), "SELECT AllergyID FROM AllergyProducts WHERE ProductID = $1", productID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var allergyID int
		if err := rows.Scan(&allergyID); err != nil {
			return nil, err
		}
		allergyIDs = append(allergyIDs, allergyID)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return allergyIDs, nil
}


func RemoveAllergyFromProduct(dbPool *pgxpool.Pool, productID, allergyID int) error {
	_, err := dbPool.Exec(context.Background(), "DELETE FROM AllergyProducts WHERE ProductID = $1 AND AllergyID = $2", productID, allergyID)
	if err != nil {
		return err
	}
	return nil
}
