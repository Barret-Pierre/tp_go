package email

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"go/tp/entities"
	"go/tp/utils"
	"mime/multipart"
	"net/smtp"
	"net/textproto"
)

func SendMail(to string, subject string, order *entities.Order, product *entities.Product, client *entities.Client) error { // Configuration de l'e-mail
	from := "from@example.com"
	body := fmt.Sprintf("Merci pour votre achat, \nvous trouverez ci joint le récapitulatif de votre commande N°%d \n\nProduit : %s \nQuantité : %d \nPrix : %.2f€ \nDate : %s \n\nCordialement, \nL'équipe de la boutique en ligne", order.ID, product.Title, order.Quantity, order.Price, order.BuyingDate.Format("2006-01-02"))

	// Configuration du serveur SMTP
	smtpHost := "localhost"
	smtpPort := "1025"
	// Génération du PDF
	pdfData, err := utils.CreatePDF(order, product, client)
	if err != nil {
		fmt.Println("Erreur lors de la création du PDF:", err)
		return err
	}

	// Construction du message MIME avec pièce jointe
	var msg bytes.Buffer
	writer := multipart.NewWriter(&msg)

	// Partie du texte
	msg.WriteString("From: " + from + "\r\n")
	msg.WriteString("To: " + to + "\r\n")
	msg.WriteString("Subject: " + subject + "\r\n")
	msg.WriteString("MIME-Version: 1.0\r\n")
	msg.WriteString(fmt.Sprintf("Content-Type: multipart/mixed; boundary=%s\r\n", writer.Boundary()))
	msg.WriteString("\r\n")

	// Partie du texte
	textPart, _ := writer.CreatePart(textproto.MIMEHeader{"Content-Type": {"text/plain; charset=utf-8"}})
	textPart.Write([]byte(body))

	// Partie du PDF
	pdfPart, _ := writer.CreatePart(textproto.MIMEHeader{
		"Content-Type":              {"application/pdf; name=\"recapitulatif.pdf\""},
		"Content-Transfer-Encoding": {"base64"},
		"Content-Disposition":       {"attachment; filename=\"recapitulatif.pdf\""},
	})

	// Encodage en base64 du PDF
	b := base64.NewEncoder(base64.StdEncoding, pdfPart)
	b.Write(pdfData)
	b.Close()

	writer.Close()

	// Envoi de l'e-mail
	err = smtp.SendMail(smtpHost+":"+smtpPort, nil, from, []string{to}, msg.Bytes())
	if err != nil {
		fmt.Println("Erreur lors de l'envoi de l'e-mail:", err)
		return err
	}

	return nil
}
