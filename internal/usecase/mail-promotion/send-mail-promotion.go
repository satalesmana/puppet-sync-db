package mailPromotion

import (
	repo "puppet-sync-db/internal/repository/send-mail"
)

func (uc *Uscase) SendMailPromotion() {
	repo := repo.NewRepoHandler(uc.config)
	repo.SendEmail("lesmanasata@gmail.com")
}
