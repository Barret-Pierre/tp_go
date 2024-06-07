package models

import (
	"database/sql"
	"fmt"
	"go/tp/entities"

	_ "github.com/go-sql-driver/mysql"
)

func CreateClient(db *sql.DB, client *entities.Client) (int64, error) {
	fmt.Println("CLIENT", client)
	result, err := db.Exec("INSERT INTO clients (firstname, lastname, address, phone, email) VALUES (?, ?, ?, ?, ?)",
		client.Firstname, client.Lastname, client.Address, client.Phone, client.Email)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func GetAllClients(db *sql.DB) ([]*entities.Client, error) {
	var clients []*entities.Client

	rows, err := db.Query("SELECT id, firstname, lastname, address, phone, email, is_active FROM clients")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var client entities.Client
		err := rows.Scan(&client.ID, &client.Firstname, &client.Lastname, &client.Address, &client.Phone, &client.Email, &client.IsActive)
		if err != nil {
			return nil, err
		}
		clients = append(clients, &client)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return clients, nil
}

func GetClientById(db *sql.DB, id int) (*entities.Client, error) {
	var client entities.Client

	row := db.QueryRow("SELECT id, firstname, lastname, address, phone, email, is_active FROM clients WHERE id = ?", id)
	err := row.Scan(&client.ID, &client.Firstname, &client.Lastname, &client.Address, &client.Phone, &client.Email, &client.IsActive)
	if err != nil {
		return nil, err
	}

	return &client, nil
}

func UpdateClient(db *sql.DB, client *entities.Client) (int64, error) {
	result, err := db.Exec("UPDATE clients SET firstname = ?, lastname = ?, address = ?, phone = ?, email = ? WHERE id = ?",
		client.Firstname, client.Lastname, client.Address, client.Phone, client.Email, client.ID)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func DeleteClient(db *sql.DB, client *entities.Client) (int64, error) {
	result, err := db.Exec("UPDATE clients SET is_active = ? WHERE id = ?", false, client.ID)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}
