package email_client

import (
	"net/smtp"
)

type EmailClientInterface interface {
	SendEmail(recieverEmail, msg string) error
}

type EmailClientImpl struct {
	AppEmail string
	Password string
}

func (e *EmailClientImpl) SendEmail(recieverEmail, msg string) error {
	auth := smtp.PlainAuth(
		"",
		e.AppEmail,
		e.Password,
		"smtp.gmail.com",
	)

	return smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		e.AppEmail,
		[]string{recieverEmail},
		[]byte(msg),
	)
}
