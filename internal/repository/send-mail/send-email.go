package sendmail

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"text/template"
	"time"

	"gopkg.in/gomail.v2"
)

func (r *Repo) SendEmail(to string, dialer *gomail.Dialer) {
	t, _ := template.ParseFiles(r.Cfg.Email.Template)
	getTime := time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local)
	time := getTime.Format("2006-01-02 15:4:5")
	subject := strings.Replace(r.Cfg.Email.Subject, "__DATE_TIME__", time, -1)

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
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", result)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		log.Fatal(err.Error())
	}
}
