package entities

import (
	"time"
)

type Order struct {
	ID         int
	IdClient   int
	IdProduct  int
	Quantity   int
	Price      float64
	BuyingDate time.Time
	Product
	Client
}

func NewOrder(idClient int, idProduct int, quantity int, price float64, buyingDate time.Time) *Order {
	return &Order{
		IdClient:   idClient,
		IdProduct:  idProduct,
		Quantity:   quantity,
		Price:      price,
		BuyingDate: buyingDate,
	}
}
