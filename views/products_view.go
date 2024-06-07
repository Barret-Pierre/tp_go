package views

import (
	"bufio"
	"fmt"
	"go/tp/entities"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func PrintMultipleProducts(products []*entities.Product) {
	if len(products) == 0 {
		fmt.Print("---- Aucun produits ----\n")
	} else {
		for _, product := range products {
			PrintProduct(product)
		}
		fmt.Print("\n")
	}
}

func PrintProduct(product *entities.Product) {
	fmt.Printf("ID: %d, Titre: %s, Description: %s, Prix: %.2f€, Quantité: %d, Is active: %v\n", product.ID, product.Title, product.Description, product.Price, product.Quantity, product.IsActive)
}

func CreateProductFromPrompt() *entities.Product {
	title := PromptProductTitle()
	description := PromptProductDescription()
	price := PromptProductPrice()
	quantity := PromptProductQuantity()
	return entities.NewProduct(title, description, price, quantity, true)
}

func PromptIdProduct(message string) int {
	var id int
	PrintMessage(message)
	fmt.Scanln(&id)
	return id
}

func PromptProductTitle() string {
	var title string
	fmt.Printf("\nEntrez le titre du produit : ")
	fmt.Scanln(&title)
	return title
}

func PromptProductDescription() string {
	var description string
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("\nEntrez la description du produit : ")
	description, _ = reader.ReadString('\n')
	description = strings.TrimSpace(description)
	return description
}

func PromptProductQuantity() int {
	var quantity int
	fmt.Printf("\nEntrez la quantité du produit : ")
	fmt.Scanln(&quantity)
	return quantity
}

func PromptProductPrice() float64 {
	var price float64
	fmt.Printf("\nEntrez le prix du produit : ")
	fmt.Scanln(&price)
	return price
}
