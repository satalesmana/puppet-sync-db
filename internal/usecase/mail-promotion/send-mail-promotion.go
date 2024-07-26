package mailPromotion

import (
	"context"
	"log"
	repoSendMail "puppet-sync-db/internal/repository/send-mail"
)

func (r *Uscase) SendMailPromotion() {
	sendMailRepo := repoSendMail.NewRepoHandler(r.config, r.connDb)

	mailToSend, err := sendMailRepo.FindEmail(context.Background(), 1)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Fatal("Send email to " + mailToSend.Email)
	// sendMailRepo.SendEmail(mailToSend[0].Email)
}
