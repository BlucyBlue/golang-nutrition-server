package db

import "database/sql"

func AddUserAllergy(db *sql.DB, userID int, allergyID int) error {
	query := `INSERT INTO UserAllergies (UserID, AllergyID) VALUES ($1, $2)`
	_, err := db.Exec(query, userID, allergyID)
	if err != nil {
		return err
	}
	return nil
}


func GetUserAllergies(db *sql.DB, userID int) ([]int, error) {
	var allergies []int

	query := `SELECT AllergyID FROM UserAllergies WHERE UserID = $1`
	rows, err := db.Query(query, userID)
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

func RemoveUserAllergy(db *sql.DB, userID int, allergyID int) error {
	query := `DELETE FROM UserAllergies WHERE UserID = $1 AND AllergyID = $2`
	_, err := db.Exec(query, userID, allergyID)
	if err != nil {
		return err
	}
	return nil
}

