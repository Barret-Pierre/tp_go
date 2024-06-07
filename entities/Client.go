package entities

import "fmt"

type Client struct {
	ID        int
	Firstname string
	Lastname  string
	Address   string
	Phone     string
	Email     string
	IsActive  bool
	TotalBill float64
}

func NewClient(firstname string, lastname string, address string, phone string, email string, isActive bool) *Client {
	return &Client{
		Firstname: firstname,
		Lastname:  lastname,
		Address:   address,
		Phone:     phone,
		Email:     email,
		IsActive:  isActive,
	}
}

func (client Client) ConvertInLine() []string {
	strIsActive := fmt.Sprintf("%v", client.IsActive)
	line := []string{client.Firstname, client.Lastname, client.Address, client.Phone, client.Email, strIsActive}
	return line
}
