package sendmail

import (
	"bytes"
	"fmt"
	"log"
	"text/template"

	"gopkg.in/gomail.v2"
)

func (r *Repo) SendEmail(to string) {
	t, _ := template.ParseFiles(r.Cfg.Email.Template)

	var body bytes.Buffer
	mimeHeaders := "\n\n"
	body.Write([]byte(fmt.Sprintf(" \n%s\n\n", mimeHeaders)))

	t.Execute(&body, struct {
		Name    string
		Message string
	}{
		Name:    "Puneet Singh",
		Message: "This is a test message in a HTML template",
	})

	result := body.String()

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", r.Cfg.Email.SenderName)
	mailer.SetHeader("To", to)
	mailer.SetHeader("Subject", r.Cfg.Email.Subject)
	mailer.SetBody("text/html", result)

	dialer := gomail.NewDialer(
		r.Cfg.Email.Host,
		r.Cfg.Email.Port,
		r.Cfg.Email.AuthEmail,
		r.Cfg.Email.AuthPassword,
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		log.Fatal(err.Error())
	}
}
