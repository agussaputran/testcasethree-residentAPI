package others

import (
	"log"
	"net/smtp"
	"os"
	"strings"
	"testcasethree-residentAPI/connection"
	"testcasethree-residentAPI/models"

	"github.com/joho/godotenv"
)

// CheckQueue func
func CheckQueue() {
	var queue models.QueueEmail
	db := connection.Connect()

	db.Where("handled = ?", false).First(&queue)
	to := []string{queue.To}
	cc := []string{queue.Cc}
	message := queue.Message
	subject := queue.Subject

	err := SendMailConfig(to, cc, subject, message)
	if err != nil {
		log.Println("Hey you have an error - > ", err.Error())
	}
	log.Println("Mail Sent")

	db.Model(&queue).Where("id = ?", queue.ID).Update("handled", true)
}

// SendMailInitial func
func SendMailInitial() {
	to := []string{"agussaputran@yopmail.com"}
	cc := []string{"gustontoi@gmail.com"}

	subject := "Test Mail"
	message := "Hello ini email uji coba"

	err := SendMailConfig(to, cc, subject, message)
	if err != nil {
		log.Println("Hey you have an error - > ", err.Error())
	}
	log.Println("Mail Sent")
}

// SendMailConfig func
func SendMailConfig(to []string, cc []string, subject, message string) error {
	err := godotenv.Load(".env")
	if err == nil {
		body := "From: " + os.Getenv("MAIL_EMAIL") + "\n" +
			"To: " + strings.Join(to, ",") + "\n" +
			"Cc: " + strings.Join(cc, ",") + "\n" +
			"Subject: " + subject + "\n\n" + message

		auth := smtp.PlainAuth("", os.Getenv("MAIL_EMAIL"), os.Getenv("MAIL_PASSWORD"), os.Getenv("MAIL_SMTP_HOST"))
		smtpAddr := os.Getenv("MAIL_SMTP_HOST") + ":" + os.Getenv("MAIL_SMTP_PORT")
		err := smtp.SendMail(smtpAddr, auth, os.Getenv("MAIL_EMAIL"), append(to, cc...), []byte(body))
		if err == nil {
			return nil
		}
		return err
	}
	return err
}
