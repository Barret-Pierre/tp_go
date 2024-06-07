package views

import (
	"bufio"
	"fmt"
	"go/tp/entities"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func PrintMultipleClients(clients []*entities.Client) {
	if len(clients) == 0 {
		fmt.Print("---- Aucun clients ----\n")
	} else {
		for _, client := range clients {
			PrintClient(client)
		}
		fmt.Print("\n")
	}
}

func PrintClient(client *entities.Client) {
	fmt.Printf("ID: %d, Prénom: %s, Nom: %s, Adresse: %s, Téléphone: %s, Email: %s, Is active: %v\n", client.ID, client.Firstname, client.Lastname, client.Address, client.Phone, client.Email, client.IsActive)
}

func CreateClientFromPrompt() *entities.Client {
	firstname := PromptClientFirstname()
	lastname := PromptClientLastname()
	phone := PromptClientPhone()
	address := PromptClientAddress()
	email := PromptClientEmail()
	return entities.NewClient(firstname, lastname, address, phone, email, true)
}

func PromptIdClient(message string) int {
	var id int
	PrintMessage(message)
	fmt.Scanln(&id)
	return id
}

func PromptClientFirstname() string {
	var firstname string
	fmt.Printf("\nEntrez le prénom du client : ")
	fmt.Scanln(&firstname)
	return firstname
}
func PromptClientLastname() string {
	var lastname string
	fmt.Printf("\nEntrez le nom du client : ")
	fmt.Scanln(&lastname)
	return lastname
}
func PromptClientPhone() string {
	var phone string
	fmt.Printf("\nEntrez le téléphone du client : ")
	fmt.Scanln(&phone)
	return phone
}
func PromptClientAddress() string {
	var address string
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("\nEntrez l'adresse du client : ")
	address, _ = reader.ReadString('\n')
	address = strings.TrimSpace(address)
	return address
}
func PromptClientEmail() string {
	var email string
	fmt.Printf("\nEntrez l'email du client : ")
	fmt.Scanln(&email)
	return email
}
