package db

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Product struct {
	ProductID int
	Name      string
	Category  string
}

func AddProduct(dbPool *pgxpool.Pool, name, category string) (int, error) {
	var productID int
	err := dbPool.QueryRow(context.Background(), "INSERT INTO Products (Name, Category) VALUES ($1, $2) RETURNING ProductID", name, category).Scan(&productID)
	if err != nil {
		return 0, err
	}
	return productID, nil
}

func GetProductByID(dbPool *pgxpool.Pool, productID int) (*Product, error) {
	product := &Product{}
	err := dbPool.QueryRow(context.Background(), "SELECT ProductID, Name, Category FROM Products WHERE ProductID = $1", productID).Scan(&product.ProductID, &product.Name, &product.Category)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func DeleteProduct(dbPool *pgxpool.Pool, productID int) error {
	_, err := dbPool.Exec(context.Background(), "DELETE FROM Products WHERE ProductID = $1", productID)
	return err
}
