package controllers

import (
	"database/sql"
	"fmt"
	"go/tp/models"
	"go/tp/utils"
	"go/tp/views"

	_ "github.com/go-sql-driver/mysql"
)

func CreateClientController(db *sql.DB) {
	client := views.CreateClientFromPrompt()
	clientId, error := models.CreateClient(db, client)
	if error != nil {
		views.PrintMessage("Erreur lors de la création du client")
	}
	views.PrintMessage(fmt.Sprintf("Client créé avec l'ID %d", clientId))
}

func ShowClientsController(db *sql.DB) {
	clients, err := models.GetAllClients(db)
	if err != nil {
		views.PrintMessage("Erreur lors de la récupération des clients")
	}
	views.PrintMessage("Tous les clients: ")
	views.PrintMultipleClients(clients)
}

func UpdateClientController(db *sql.DB) {
	clientId := views.PromptIdClient("Entrez l'id du client à modifier: ")
	client, getClientError := models.GetClientById(db, clientId)
	if getClientError != nil {
		views.PrintMessage("Ce client n'existe pas")
		return
	}

	views.PrintMessage("Voici le client à modifier:")
	views.PrintClient(client)
	updateNavItems := []string{"Prénom", "Nom", "Address", "Phone", "Email", "Quitter"}
	views.ShowUpdateMenu(updateNavItems)

	var choice int
	fmt.Scanln(&choice)
	for choice != 6 {
		switch choice {
		case 1:
			client.Firstname = views.PromptClientFirstname()
		case 2:
			client.Lastname = views.PromptClientLastname()
		case 3:
			client.Address = views.PromptClientAddress()
		case 4:
			client.Phone = views.PromptClientPhone()
		case 5:
			client.Email = views.PromptClientEmail()
		default:
			views.PrintMessage("Choix invalide")
		}

		_, error := models.UpdateClient(db, client)
		if error != nil {
			views.PrintMessage("Erreur lors de la mise à jour du client")
			return
		}
		views.PrintMessage("Client mis à jour")

		client, _ = models.GetClientById(db, clientId)
		views.PrintMessage("Voici le client à modifier:")
		views.PrintClient(client)
		views.ShowUpdateMenu(updateNavItems)
		fmt.Scanln(&choice)
	}
}

func DeleteClientController(db *sql.DB) {
	clientId := views.PromptIdClient("Entrez l'id du client à supprimer: ")
	client, getClientError := models.GetClientById(db, clientId)
	if getClientError != nil {
		views.PrintMessage("Ce client n'existe pas")
		return
	}

	client.IsActive = false

	_, error := models.DeleteClient(db, client)
	if error != nil {
		views.PrintMessage("Erreur lors de la suppresion du client")
		return
	}
	views.PrintMessage("Client supprimé !")
}

func ExportClientsControllerToCSV(db *sql.DB) {
	clients, err := models.GetAllClients(db)
	if err != nil {
		views.PrintMessage("Erreur lors de la récupération des clients")
		return
	}
	utils.CreateCSVClientsFile("./exports/clients.csv", clients)
	views.PrintMessage("Clients exportés avec succès !")
}
