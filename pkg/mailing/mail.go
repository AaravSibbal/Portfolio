package mailing

import (
	"fmt"
	"net/smtp"
)

func Mail(name, clientEmail, phone, discription, bEmail, bPassword string) error {
	auth := smtp.PlainAuth(
		"",
		bEmail,
		bPassword,
		"smtp.gmail.com",
	)

	msg := fmt.Sprintf("Subject: Booking inquiry by %s Phone: %s Email:%s\n%s", name, phone, clientEmail, discription)

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		bEmail,
		[]string{bEmail},
		[]byte(msg),
	)

	return err
}