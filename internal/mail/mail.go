package mail

import (
	"bytes"
	"html/template"

	"gopkg.in/gomail.v2"
)

type EmailData struct {
	Name  string
	Email string
}

func SendHTMLEmail(to string, subject string, data EmailData) error {
	// Parse HTML file
	tmpl, err := template.ParseFiles("internal/mail/template/forgot.html")
	if err != nil {
		return err
	}

	var body bytes.Buffer
	if err := tmpl.Execute(&body, data); err != nil {
		return err
	}

	// Compose email
	m := gomail.NewMessage()
	m.SetHeader("From", "your-email@gmail.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body.String())

	// SMTP config
	d := gomail.NewDialer("sandbox.smtp.mailtrap.io", 587, "90b2b0e2fcc82f", "8e1fa4ffabe213")

	return d.DialAndSend(m)
}
