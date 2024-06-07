package views

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ShowMenu(navItems []string) {
	PrintSeparator()
	fmt.Printf("Que voulez-vous faire?\n")
	for i, item := range navItems {
		if i+1 == 1 {
			PrintMessage("---Products---")
		}
		if i+1 == 6 {
			PrintMessage("---Clients---")
		}
		if i+1 == 11 {
			PrintMessage("---Commandes---")
		}
		fmt.Printf("%d- %s\n", i+1, item)
	}
	PrintSeparator()
}

func ShowUpdateMenu(navItems []string) {
	PrintSeparator()
	fmt.Printf("Quel champs voulez vous mettre à jour?\n")
	for i, item := range navItems {
		fmt.Printf("%d- %s\n", i+1, item)
	}
	PrintSeparator()
}

func PrintMessage(message string) {
	fmt.Print("\n", message, "\n")
}

func PrintSeparator() {
	fmt.Print("\n##############################################\n\n")
}
