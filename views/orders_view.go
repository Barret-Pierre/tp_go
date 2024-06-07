package views

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func PromptOrderQuantity() int {
	var quantity int
	fmt.Printf("\nEntrez la quantité que vous souhaitez commander : \n")
	fmt.Scanln(&quantity)
	return quantity
}
