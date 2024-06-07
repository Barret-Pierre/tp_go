package controllers

import (
	"database/sql"
	"fmt"
	"go/tp/models"
	"go/tp/utils"
	"go/tp/views"

	_ "github.com/go-sql-driver/mysql"
)

func CreateProductController(db *sql.DB) {
	product := views.CreateProductFromPrompt()
	productId, error := models.CreateProduct(db, product)
	if error != nil {
		views.PrintMessage("Erreur lors de la création du produit")
	}
	views.PrintMessage(fmt.Sprintf("Produit créé avec l'ID %d", productId))
}

func ShowProductsController(db *sql.DB) {
	products, err := models.GetAllProducts(db)
	if err != nil {
		views.PrintMessage("Erreur lors de la récupération des produits")
	}
	views.PrintMessage("Tous les produits: \n")
	views.PrintMultipleProducts(products)
}

func UpdateProductController(db *sql.DB) {
	productId := views.PromptIdProduct("Entrez l'id du produit à modifier: ")
	product, getProductError := models.GetProductById(db, productId)
	if getProductError != nil {
		views.PrintMessage("Ce produit n'existe pas")
		return
	}

	views.PrintMessage("Voici le produit à modifier:")
	views.PrintProduct(product)
	updateNavItems := []string{"Titre", "Description", "Quantité", "Prix", "Quitter"}
	views.ShowUpdateMenu(updateNavItems)

	var choice int
	fmt.Scanln(&choice)
	for choice != 5 {
		switch choice {
		case 1:
			product.Title = views.PromptProductTitle()
		case 2:
			product.Description = views.PromptProductDescription()
		case 3:
			product.Quantity = views.PromptProductQuantity()
		case 4:
			product.Price = views.PromptProductPrice()
		default:
			views.PrintMessage("Choix invalide")
		}

		_, error := models.UpdateProduct(db, product)
		if error != nil {
			views.PrintMessage("Erreur lors de la mise à jour du produit")
			return
		}
		views.PrintMessage("Produit mis à jour")

		product, _ = models.GetProductById(db, productId)
		views.PrintMessage("Voici le produit à modifier:")
		views.PrintProduct(product)
		views.ShowUpdateMenu(updateNavItems)
		fmt.Scanln(&choice)
	}
}

func DeleteProductController(db *sql.DB) {
	productId := views.PromptIdProduct("Entrez l'id du produit à supprimer: ")
	product, getProductError := models.GetProductById(db, productId)
	if getProductError != nil {
		views.PrintMessage("Ce produit n'existe pas")
		return
	}

	product.IsActive = false

	_, error := models.DeleteProduct(db, product)
	if error != nil {
		views.PrintMessage("Erreur lors de la suppresion du produit")
		return
	}
	views.PrintMessage("Produit supprimé !")
}

func ExportProductsControllerToCSV(db *sql.DB) {
	products, err := models.GetAllProducts(db)
	if err != nil {
		views.PrintMessage("Erreur lors de la récupération des produits")
		return
	}
	utils.CreateCSVProductsFile("./exports/produits.csv", products)
	views.PrintMessage("Produits exportés avec succès !")
}
