package syncdb

import (
	model "puppet-sync-db/internal/model/config"

	"go.mongodb.org/mongo-driver/mongo"
)

type Uscase struct {
	config          *model.Config
	mongoConnRemote *mongo.Database
	mongoConnLocal  *mongo.Database
}

func NewSyncDBHandler(config *model.Config, conDbRemote *mongo.Database, conDbLocal *mongo.Database) Handler {
	return &Uscase{
		config:          config,
		mongoConnRemote: conDbRemote,
		mongoConnLocal:  conDbLocal,
	}
}

type Handler interface {
	SyncDBToLocal()
}
