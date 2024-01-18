package jobstereet_pelamar

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type Repo struct {
	Db *mongo.Database
}

func NewRepoHandler(conDbRemote *mongo.Database) Handler {
	return &Repo{
		Db: conDbRemote,
	}
}

type Handler interface {
	FetchListData(ctx context.Context, limit int64) ([]interface{}, error)
}
