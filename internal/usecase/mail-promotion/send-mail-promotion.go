package mailPromotion

import (
	"context"
	"log"
	repoSendMail "puppet-sync-db/internal/repository/send-mail"
	"strconv"

	"gopkg.in/gomail.v2"
)

func (r *Uscase) SendMailPromotion() {
	dialer := gomail.NewDialer(
		r.config.Email.Host,
		r.config.Email.Port,
		r.config.Email.AuthEmail,
		r.config.Email.AuthPassword,
	)

	limit, err := strconv.ParseInt(r.config.Email.Limit, 10, 64)
	if err != nil {
		log.Fatal(err.Error())
	}

	sendMailRepo := repoSendMail.NewRepoHandler(r.config, r.connDb)
	sendMailRepo.SendEmail("lesmanasata@gmail.com", dialer)
	mailToSend, err := sendMailRepo.FindEmail(context.Background(), limit)
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, activity := range mailToSend {
		log.Printf("Send email to " + activity.Email)
		sendMailRepo.SendEmail(activity.Email, dialer)

		_, errFlag := sendMailRepo.SetFlagMail(context.Background(), activity.Email)
		if errFlag != nil {
			log.Fatal(err.Error())
		}
	}
}
