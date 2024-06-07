package entities

import (
	"fmt"
	"strconv"
)

type Product struct {
	ID          int
	Title       string
	Description string
	Quantity    int
	Price       float64
	IsActive    bool
}

func NewProduct(title string, description string, price float64, quantity int, isActive bool) *Product {
	return &Product{
		Title:       title,
		Description: description,
		Price:       price,
		Quantity:    quantity,
		IsActive:    isActive,
	}
}

func (product Product) ConvertInLine() []string {
	strPrice := strconv.FormatFloat(product.Price, 'f', 2, 64)
	strQuantity := strconv.Itoa(product.Quantity)
	strIsActive := fmt.Sprintf("%v", product.IsActive)
	line := []string{product.Title, product.Description, strPrice, strQuantity, strIsActive}
	return line
}
