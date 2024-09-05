package mailPromotion

import (
	model "puppet-sync-db/internal/model/config"

	"go.mongodb.org/mongo-driver/mongo"
)

type Uscase struct {
	config      *model.Config
	connDb      *mongo.Database
	conDBRemote *mongo.Database
}

func NewMailPromotionHandler(config *model.Config, coonDb *mongo.Database, conDBRemote *mongo.Database) Handler {
	return &Uscase{
		config:      config,
		connDb:      coonDb,
		conDBRemote: conDBRemote,
	}
}

type Handler interface {
	SendMailPromotion()
}
