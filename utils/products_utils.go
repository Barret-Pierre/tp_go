package utils

import (
	"encoding/csv"
	"go/tp/entities"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func CreateCSVProductsFile(filePath string, products []*entities.Product) {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("Erreur lors de la lecture du fichier: %v", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, product := range products {
		errorWritter := writer.Write(product.ConvertInLine())
		if errorWritter != nil {
			log.Fatalf("Erreur lors de la lecture du fichier: %v", err)
			return
		}
	}
}
