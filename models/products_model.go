package models

import (
	"database/sql"
	"go/tp/entities"

	_ "github.com/go-sql-driver/mysql"
)

func CreateProduct(db *sql.DB, product *entities.Product) (int64, error) {
	result, err := db.Exec("INSERT INTO products (title, description, quantity, price) VALUES (?, ?, ?, ?)",
		product.Title, product.Description, product.Quantity, product.Price)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func GetAllProducts(db *sql.DB) ([]*entities.Product, error) {
	var products []*entities.Product

	rows, err := db.Query("SELECT id, title, description, quantity, price, is_active FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product entities.Product
		err := rows.Scan(&product.ID, &product.Title, &product.Description, &product.Quantity, &product.Price, &product.IsActive)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func GetProductById(db *sql.DB, id int) (*entities.Product, error) {
	var product entities.Product

	row := db.QueryRow("SELECT id, title, description, quantity, price, is_active FROM products WHERE id = ?", id)
	err := row.Scan(&product.ID, &product.Title, &product.Description, &product.Quantity, &product.Price, &product.IsActive)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func UpdateProduct(db *sql.DB, product *entities.Product) (int64, error) {
	result, err := db.Exec("UPDATE products SET title = ?, description = ?, quantity = ?, price = ? WHERE id = ?",
		product.Title, product.Description, product.Quantity, product.Price, product.ID)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func DeleteProduct(db *sql.DB, product *entities.Product) (int64, error) {
	result, err := db.Exec("UPDATE products SET is_active = ? WHERE id = ?", false, product.ID)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}
