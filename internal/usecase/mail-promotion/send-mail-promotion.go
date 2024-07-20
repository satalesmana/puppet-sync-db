package mailPromotion

import (
	repoSendMail "puppet-sync-db/internal/repository/send-mail"
)

func (r *Uscase) SendMailPromotion() {
	sendMailRepo := repoSendMail.NewRepoHandler(r.config)
	sendMailRepo.SendEmail("lesmanasata@gmail.com")
}
