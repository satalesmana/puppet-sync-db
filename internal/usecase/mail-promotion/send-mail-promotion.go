package mailPromotion

import (
	"context"
	"log"
	model "puppet-sync-db/internal/model/local-pelamars"
	repoSendMail "puppet-sync-db/internal/repository/send-mail"
)

func (r *Uscase) SendMailPromotion() {
	var mailToSend *model.Activity
	sendMailRepo := repoSendMail.NewRepoHandler(r.config, r.connDb)

	mailToSend, err := sendMailRepo.FindEmail(context.Background())
	if err != nil {
		log.Fatal(err.Error())
	}

	sendMailRepo.SendEmail(mailToSend.Email)
}
