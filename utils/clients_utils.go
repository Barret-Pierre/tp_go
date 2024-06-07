package utils

import (
	"encoding/csv"
	"go/tp/entities"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func CreateCSVClientsFile(filePath string, clients []*entities.Client) {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("Erreur lors de la lecture du fichier: %v", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, client := range clients {
		errorWritter := writer.Write(client.ConvertInLine())
		if errorWritter != nil {
			log.Fatalf("Erreur lors de la lecture du fichier: %v", err)
			return
		}
	}
}
