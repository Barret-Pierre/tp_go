package controllers

import (
	"database/sql"
	"fmt"
	"go/tp/email"
	"go/tp/entities"
	"go/tp/models"
	"go/tp/utils"
	"go/tp/views"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func CreateOrderController(db *sql.DB) {

	clientId := views.PromptIdClient("Entrez l'ID du client qui commande")
	client, clientError := models.GetClientById(db, clientId)
	if clientError != nil {
		views.PrintMessage("Ce client n'existe pas !")
		return
	}
	if !client.IsActive {
		views.PrintMessage("Ce client n'existe plus !")
		return
	}

	productId := views.PromptIdClient("Entrez l'ID du produit à commander")
	product, productError := models.GetProductById(db, productId)
	if productError != nil {
		views.PrintMessage("Ce produit n'existe pas !")
		return
	}
	if !product.IsActive || product.Quantity == 0 {
		views.PrintMessage("Ce produit n'est plus disponible !")
		return
	}

	quantity := views.PromptOrderQuantity()

	for quantity > product.Quantity {
		if quantity > product.Quantity {
			views.PrintMessage(fmt.Sprintf("Pas assez de quantité disponible, quantité maximale : %d", product.Quantity))
		}
		quantity = views.PromptOrderQuantity()
	}

	price := product.Price * float64(quantity)
	currentDate := time.Now()
	order := entities.NewOrder(clientId, productId, quantity, price, currentDate)
	orderId, error := models.CreateOrder(db, order)
	if error != nil {
		fmt.Print(error)
		views.PrintMessage("Erreur lors de la création de la commande !")
		return
	}
	order.ID = int(orderId)

	product.Quantity -= quantity
	_, updateError := models.UpdateProduct(db, product)
	if updateError != nil {
		views.PrintMessage("Erreur lors de la mise à jour des quantité du produit !")
		return
	}
	views.PrintMessage(fmt.Sprintf("Commande N°%d créée avec succès", orderId))
	views.PrintMessage(fmt.Sprintf("Vous receverez un récapitulatif par mail à l'adresse %s", client.Email))
	mailError := email.SendMail(client.Email, fmt.Sprintf("Commande N°%d", orderId), order, product, client)
	if mailError != nil {
		views.PrintMessage("Erreur lors de l'envoi du mail !")
		return
	}
	views.PrintMessage("Mail envoyé avec succès !")

}

func ExportAllCommandsBills(db *sql.DB) {
	clientWithTotalBill, err := models.GetAllSumPriceOrderGroupeByClient(db)
	if err != nil {
		views.PrintMessage("Erreur lors de la récupération des clients")
		return
	}
	ordersWithClientsAndProducts, err := models.GetAllOrderWithClientAndProduct(db)
	if err != nil {
		views.PrintMessage("Erreur lors de la récupération des commandes")
		return
	}

	buildPdfError := utils.CreateTotalBillPDF(ordersWithClientsAndProducts, clientWithTotalBill)
	if buildPdfError != nil {
		views.PrintMessage("Erreur lors de la création du pdf")
		return
	}

	views.PrintMessage("Total des commandes exporté avec succès !")
}
