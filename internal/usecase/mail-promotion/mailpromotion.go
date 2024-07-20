package mailPromotion

import (
	model "puppet-sync-db/internal/model/config"

	"go.mongodb.org/mongo-driver/mongo"
)

type Uscase struct {
	config *model.Config
	connDb *mongo.Database
}

func NewMailPromotionHandler(config *model.Config, coonDb *mongo.Database) Handler {
	return &Uscase{
		config: config,
		connDb: coonDb,
	}
}

type Handler interface {
	SendMailPromotion()
}
