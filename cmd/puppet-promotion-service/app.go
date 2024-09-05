package main

import (
	"log"
	model "puppet-sync-db/internal/model/config"
	mailPromotion "puppet-sync-db/internal/usecase/mail-promotion"
)

func startApp(config *model.Config) error {

	connDb, err := setupMongoDBConnection(config, "local")
	if err != nil {
		log.Fatalf("error startup service when connection setup to mongodb. err: %v", err)
		return err
	}

	connDBRemote, err := setupMongoDBConnection(config, "remote")
	if err != nil {
		log.Fatalf("error startup service when connection setup to mongodb. err: %v", err)
		return err
	}

	sendMail := mailPromotion.NewMailPromotionHandler(config, connDb, connDBRemote)
	sendMail.SendMailPromotion()

	log.Println("Successfully send")
	return nil
}
