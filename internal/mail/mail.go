package mail

import (
	"bytes"
	"html/template"
	"strconv"

	"github.com/lasamarndi1994/gov2/internal/config"
	"gopkg.in/gomail.v2"
)

type EmailData struct {
	Name       string
	Email      string
	ResetToken string
}

var cfg = config.LoadConfig()

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
	port, _ := strconv.Atoi(cfg.MailPort)
	d := gomail.NewDialer(cfg.MailHost, port, cfg.MailUsername, cfg.MailPassword)

	return d.DialAndSend(m)
}
