package jobstereet_pelamar

import (
	"context"
	config "puppet-sync-db/internal/model/config"

	"go.mongodb.org/mongo-driver/mongo"
)

type Repo struct {
	cfg *config.Config
	Db  *mongo.Database
}

func NewRepoHandler(config *config.Config, conDbRemote *mongo.Database) Handler {
	return &Repo{
		cfg: config,
		Db:  conDbRemote,
	}
}

type Handler interface {
	FetchListData(ctx context.Context, limit int64) ([]interface{}, error)
}
