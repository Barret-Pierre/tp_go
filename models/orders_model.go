package models

import (
	"database/sql"
	"go/tp/entities"
	"go/tp/utils"

	_ "github.com/go-sql-driver/mysql"
)

func CreateOrder(db *sql.DB, order *entities.Order) (int64, error) {
	result, err := db.Exec("INSERT INTO orders (client_id, product_id, quantity, price, created_at) VALUES (?, ?, ?, ?, ?)",
		order.IdClient, order.IdProduct, order.Quantity, order.Price, order.BuyingDate)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func GetAllSumPriceOrderGroupeByClient(db *sql.DB) ([]*entities.Client, error) {
	var clients []*entities.Client

	rows, err := db.Query("SELECT c.id, c.firstname, c.lastname, c.email, c.address, c.phone, sum(o.price) AS total_spent FROM tp_go.clients AS c LEFT JOIN tp_go.orders AS o ON c.id = o.client_id GROUP BY c.id;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var client entities.Client
		var totalSpent sql.NullFloat64

		err := rows.Scan(&client.ID, &client.Firstname, &client.Lastname, &client.Email, &client.Address, &client.Phone, &totalSpent)
		if err != nil {
			return nil, err
		}

		if totalSpent.Valid {
			client.TotalBill = totalSpent.Float64
		} else {
			// total_spent est NULL, donc vous pouvez définir TotalBill à une valeur par défaut
			client.TotalBill = 0
		}
		clients = append(clients, &client)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return clients, nil
}

func GetAllOrderWithClientAndProduct(db *sql.DB) ([]*entities.Order, error) {
	var orders []*entities.Order

	rows, err := db.Query("SELECT o.id, o.client_id, o.product_id, o.quantity, o.price, o.created_at, c.firstname, c.lastname, c.email, c.phone, c.address, c.is_active, p.title, p.description, p.quantity, p.price, p.is_active FROM orders AS o INNER JOIN clients AS c ON o.client_id=c.id INNER JOIN products AS p ON o.product_id=p.id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var order entities.Order
		var buyingDateByte []uint8
		err := rows.Scan(&order.ID, &order.IdClient, &order.IdProduct, &order.Quantity, &order.Price, &buyingDateByte, &order.Client.Firstname, &order.Client.Lastname, &order.Client.Email, &order.Client.Phone, &order.Client.Address, &order.Client.IsActive, &order.Product.Title, &order.Product.Description, &order.Product.Quantity, &order.Product.Price, &order.Product.IsActive)
		if err != nil {
			return nil, err
		}
		order.BuyingDate = *utils.ConvertUint8ToTime(buyingDateByte)
		orders = append(orders, &order)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}

func GetOrderById(db *sql.DB, id int) (*entities.Order, error) {
	var order entities.Order

	row := db.QueryRow("SELECT id, client_id, product_id, quantity, price, created_at FROM orders WHERE id = ?", id)
	err := row.Scan(&order.ID, &order.IdClient, &order.IdProduct, &order.Quantity, &order.Price, &order.BuyingDate)
	if err != nil {
		return nil, err
	}

	return &order, nil
}
