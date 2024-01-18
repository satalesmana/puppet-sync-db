package syncdb

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Uscase struct {
	mongoConnRemote *mongo.Database
	mongoConnLocal  *mongo.Database
}

func NewSyncDBHandler(conDbRemote *mongo.Database, conDbLocal *mongo.Database) Handler {
	return &Uscase{
		mongoConnRemote: conDbRemote,
		mongoConnLocal:  conDbLocal,
	}
}

type Handler interface {
	SyncDBToLocal()
}
