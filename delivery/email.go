package delivery

import (
	"fmt"
	"net/smtp"
)

func SendEmail(to string, body string) error {
	from := "from@gmail.com"
	password := "password"
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	toArray := []string{to}

	message := []byte(body)

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, toArray, message)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("Email Sent Successfully!")
	return nil
}
