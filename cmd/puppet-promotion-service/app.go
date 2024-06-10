package main

import (
	model "puppet-sync-db/internal/model/config"
)

func startApp(config *model.Config) error {

	// connDb, err := setupMongoDBConnection(config, "local")
	// if err != nil {
	// 	log.Fatalf("error startup service when connection setup to mongodb. err: %v", err)
	// 	return err
	// }

	// sendMail := mailPromotion.NewMailPromotionHandler(config, connDb)
	// sendMail.SendMailPromotion()

	return nil
}
