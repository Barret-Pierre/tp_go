package main

import (
	"database/sql"
	"fmt"
	"go/tp/controllers"
	"go/tp/views"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Informations de connexion à la base de données
	dsn := "root:SMack9911!DOwn++@tcp(127.0.0.1:3306)/tp_go"

	// Ouvrir une connexion à la base de données
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Erreur lors de l'ouverture de la base de données: %v", err)
	}
	defer db.Close()

	// Vérifier la connexion
	err = db.Ping()
	if err != nil {
		log.Fatalf("Erreur lors de la vérification de la connexion: %v", err)
	}

	views.PrintMessage("Connexion réussie à la base de données!")

	// Initialiser la base de données en créant la table products si elle n'existe pas
	err = initializeDatabase(db)
	if err != nil {
		log.Fatalf("Erreur lors de l'initialisation de la base de données: %v", err)
	}

	navItems := []string{"Ajouter un produit", "Afficher l'ensemble des produits", "Modifier un produit", "Supprimer un produit", "Exporter l'ensemble des produits sous forme csv", "Ajouter un client", "Afficher l'ensemble des clients", "Modifier un client", "Supprimer un client", "Exporter l'ensemble des clients sous forme csv", "Effectuer une commande", "Exports l'ensemble des commande", "Quitter"}
	views.ShowMenu(navItems)

	var choice int
	fmt.Scanln(&choice)
	for choice != 13 {
		switch choice {
		case 1:
			controllers.CreateProductController(db)
		case 2:
			controllers.ShowProductsController(db)
		case 3:
			controllers.UpdateProductController(db)
		case 4:
			controllers.DeleteProductController(db)
		case 5:
			controllers.ExportProductsControllerToCSV(db)
		case 6:
			controllers.CreateClientController(db)
		case 7:
			controllers.ShowClientsController(db)
		case 8:
			controllers.UpdateClientController(db)
		case 9:
			controllers.DeleteClientController(db)
		case 10:
			controllers.ExportClientsControllerToCSV(db)
		case 11:
			controllers.CreateOrderController(db)
		case 12:
			controllers.ExportAllCommandsBills(db)
		default:
			views.PrintMessage("Choix invalide")
		}
		views.ShowMenu(navItems)
		fmt.Scanln(&choice)
	}

}

// DATABASE
func initializeDatabase(db *sql.DB) error {
	// Vérifier si la table products existe déjà
	productRows, err := db.Query("SHOW TABLES LIKE 'products'")
	if err != nil {
		return err
	}
	defer productRows.Close()

	if !productRows.Next() {
		// La table products n'existe pas, la créer
		_, err := db.Exec(`CREATE TABLE products (
            id INT AUTO_INCREMENT PRIMARY KEY,
            title VARCHAR(255),
            description TEXT,
            quantity INT,
            price DOUBLE,
            is_active BOOLEAN DEFAULT TRUE
        )`)
		if err != nil {
			return err
		}
		fmt.Println("Table 'products' créée avec succès.")
	} else {
		fmt.Println("La table 'products' existe déjà.")
	}

	// Vérifier si la table products existe déjà
	clientsRows, err := db.Query("SHOW TABLES LIKE 'clients'")
	if err != nil {
		return err
	}
	defer clientsRows.Close()

	if !clientsRows.Next() {
		// La table products n'existe pas, la créer
		_, err := db.Exec(`CREATE TABLE clients (
            id INT AUTO_INCREMENT PRIMARY KEY,
            firstname VARCHAR(255),
            lastname VARCHAR(255),
            phone VARCHAR(255),
            address VARCHAR(255),
            email VARCHAR(255),
						is_active BOOLEAN DEFAULT TRUE
        )`)
		if err != nil {
			return err
		}
		fmt.Println("Table 'clients' créée avec succès.")
	} else {
		fmt.Println("La table 'clients' existe déjà.")
	}

	// Vérifier si la table products existe déjà
	ordersRows, err := db.Query("SHOW TABLES LIKE 'orders'")
	if err != nil {
		return err
	}
	defer ordersRows.Close()

	if !ordersRows.Next() {
		// La table products n'existe pas, la créer
		_, err := db.Exec(`CREATE TABLE orders (
						id INT AUTO_INCREMENT PRIMARY KEY,
						client_id INT,
						product_id INT,
						quantity INT,
						price DOUBLE,
						created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
						FOREIGN KEY (client_id) REFERENCES clients(id),
						FOREIGN KEY (product_id) REFERENCES products(id)
        )`)
		if err != nil {
			return err
		}
		fmt.Println("Table 'orders' créée avec succès.")
	} else {
		fmt.Println("La table 'orders' existe déjà.")
	}
	return nil
}
