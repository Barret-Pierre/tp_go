package utils

import (
	"bytes"
	"fmt"
	"go/tp/entities"
	"os"
	"time"

	"github.com/jung-kurt/gofpdf"
)

func CreatePDF(order *entities.Order, product *entities.Product, client *entities.Client) ([]byte, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.AddUTF8Font("DejaVu", "", "fonts/DejaVuSans.ttf")

	pdf.SetFont("DejaVu", "", 20)
	pdf.Cell(40, 10, "Récapitulatif de commande")

	pdf.Ln(12)
	pdf.SetFont("DejaVu", "", 16)
	pdf.Cell(40, 10, "Information sur le client")
	pdf.Ln(8)
	pdf.SetFont("DejaVu", "", 12)
	pdf.Cell(40, 10, fmt.Sprintf("Riason social: %s %s", client.Lastname, client.Firstname))
	pdf.Ln(4)
	pdf.Cell(40, 10, fmt.Sprintf("Adresse : %s", client.Address))
	pdf.Ln(4)
	pdf.Cell(40, 10, fmt.Sprintf("Téléphone : %s", client.Phone))
	pdf.Ln(4)
	pdf.Cell(40, 10, fmt.Sprintf("Email : %s", client.Email))

	pdf.Ln(12)
	pdf.SetFont("DejaVu", "", 16)
	pdf.Cell(40, 10, "Information sur la commande")
	pdf.Ln(8)
	pdf.SetFont("DejaVu", "", 12)
	pdf.Cell(40, 10, fmt.Sprintf("Commande N°: %d", order.ID))
	pdf.Ln(4)
	pdf.Cell(40, 10, fmt.Sprintf("Produit : %s", product.Title))
	pdf.Ln(4)
	pdf.Cell(40, 10, fmt.Sprintf("Quantité : %d", order.Quantity))
	pdf.Ln(4)
	pdf.Cell(40, 10, fmt.Sprintf("Prix : %.2f€", order.Price))
	pdf.Ln(4)
	pdf.Cell(40, 10, fmt.Sprintf("Date : %s", order.BuyingDate.Format("2006-01-02 15:04:05")))

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func CreateTotalBillPDF(ordersWithClientsAndProducts []*entities.Order, clientsWithTotalBill []*entities.Client) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddUTF8Font("DejaVu", "", "fonts/DejaVuSans.ttf")

	for _, client := range clientsWithTotalBill {
		pdf.AddPage()
		pdf.SetFont("DejaVu", "", 20)
		pdf.Cell(40, 10, fmt.Sprintf("Client N°: %d", client.ID))
		pdf.Ln(12)
		pdf.SetFont("DejaVu", "", 16)
		pdf.Cell(40, 10, "Information sur le client")
		pdf.Ln(8)
		pdf.SetFont("DejaVu", "", 12)
		pdf.Cell(40, 10, fmt.Sprintf("Riason social: %s %s", client.Lastname, client.Firstname))
		pdf.Ln(4)
		pdf.Cell(40, 10, fmt.Sprintf("Adresse : %s", client.Address))
		pdf.Ln(4)
		pdf.Cell(40, 10, fmt.Sprintf("Téléphone : %s", client.Phone))
		pdf.Ln(4)
		pdf.Cell(40, 10, fmt.Sprintf("Email : %s", client.Email))

		pdf.Ln(12)
		pdf.SetFont("DejaVu", "", 16)
		pdf.Cell(40, 10, "Commandes passés par le client")
		clientOrder := filterOrdersByClient(ordersWithClientsAndProducts, client.ID)

		if len(clientOrder) == 0 {
			pdf.Ln(8)
			pdf.SetFont("DejaVu", "", 12)
			pdf.Cell(40, 10, "Aucune commande faite par ce client")
		} else {
			for _, order := range clientOrder {
				if order.IdClient == client.ID {
					pdf.Ln(8)
					pdf.SetFont("DejaVu", "", 12)
					pdf.Cell(40, 10, fmt.Sprintf("Commande N°: %d", order.ID))
					pdf.Ln(4)
					pdf.Cell(40, 10, fmt.Sprintf("Produit : %s", order.Product.Title))
					pdf.Ln(4)
					pdf.Cell(40, 10, fmt.Sprintf("Quantité : %d", order.Quantity))
					pdf.Ln(4)
					pdf.Cell(40, 10, fmt.Sprintf("Prix : %.2f€", order.Price))
					pdf.Ln(4)
					pdf.Cell(40, 10, fmt.Sprintf("Date : %s", order.BuyingDate.Format("2006-01-02 15:04:05")))
				}
			}

			pdf.Ln(12)
			pdf.SetFont("DejaVu", "", 16)
			pdf.Cell(40, 10, fmt.Sprintf("Prix total des commande: %.2f€", client.TotalBill))
		}
	}

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		return err
	}

	f, err := os.OpenFile("./exports/total_bill.pdf", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(buf.Bytes())
	if err != nil {
		return err
	}

	return nil

}

func filterOrdersByClient(ordersWithClientsAndProducts []*entities.Order, clientId int) []*entities.Order {
	var ordersFilter []*entities.Order
	for _, order := range ordersWithClientsAndProducts {
		if order.IdClient == clientId {
			ordersFilter = append(ordersFilter, order)
		}
	}
	return ordersFilter
}

func ConvertUint8ToTime(value []uint8) *time.Time {
	strValue := string(value)

	time, err := time.Parse("2006-01-02 15:04:05", strValue)
	if err != nil {
		fmt.Println("Erreur lors de la conversion de la chaîne en time.Time :", err)
		return nil
	}

	return &time
}
