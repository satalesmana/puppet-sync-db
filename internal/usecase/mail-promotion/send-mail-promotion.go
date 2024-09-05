package mailPromotion

import (
	"context"
	"log"
	repoSmtpAccount "puppet-sync-db/internal/repository/remote-smtpaccount"
	repoSendMail "puppet-sync-db/internal/repository/send-mail"
	"strconv"

	"gopkg.in/gomail.v2"
)

func (r *Uscase) SendMailPromotion() {
	smtpAccountRepo := repoSmtpAccount.NewRepoHandler(r.conDBRemote)
	smtpAccount, err := smtpAccountRepo.GetSmtpAccount(context.TODO())
	if err != nil {
		log.Fatal(err.Error())
	}

	dialer := gomail.NewDialer(
		smtpAccount.HostName,
		smtpAccount.Port,
		smtpAccount.AuthEmail,
		smtpAccount.AuthPassword,
	)

	senderName := "Tanya Dosen Official <" + smtpAccount.AuthEmail + ">"

	sendMailRepo := repoSendMail.NewRepoHandler(r.config, r.connDb)
	// sendMailRepo.SendEmail(dialer, "lesmanasata@gmail.com", "tes", senderName)

	limit, err := strconv.ParseInt(r.config.Email.Limit, 10, 64)
	if err != nil {
		log.Fatal(err.Error())
	}

	mailToSend, err := sendMailRepo.FindEmail(context.Background(), limit)
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, activity := range mailToSend {
		log.Printf("Send email to " + activity.Email)
		sendMailRepo.SendEmail(dialer, activity.Email, activity.Name, senderName)

		_, errFlag := sendMailRepo.SetFlagMail(context.Background(), activity.Email)
		if errFlag != nil {
			log.Fatal(err.Error())
		}
	}
}
