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

	log.Printf("Send email to " + mailToSend.Email)
	sendMailRepo.SendEmail(mailToSend.Email)

	_, errFlag := sendMailRepo.SetFlagMail(context.Background(), mailToSend.Email)
	if errFlag != nil {
		log.Fatal(err.Error())
	}

}
