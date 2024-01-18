package local_pelamar

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type Repo struct {
	Db *mongo.Database
}

func NewRepoHandler(conDbLocal *mongo.Database) Handler {
	return &Repo{
		Db: conDbLocal,
	}
}

type Handler interface {
	StoreMultiplePelamar(ctx context.Context, data []interface{}) (int64, error)
	StoreMultipleActivitys(ctx context.Context, data []interface{}) (int64, error)
}
