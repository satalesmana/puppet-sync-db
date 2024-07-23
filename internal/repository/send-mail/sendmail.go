package sendmail

import (
	"context"
	config "puppet-sync-db/internal/model/config"
	model "puppet-sync-db/internal/model/local-pelamars"

	"go.mongodb.org/mongo-driver/mongo"
)

type Repo struct {
	Cfg *config.Config
	Db  *mongo.Database
}

func NewRepoHandler(cfg *config.Config, Db *mongo.Database) Handler {
	return &Repo{
		Cfg: cfg,
		Db:  Db,
	}
}

type Handler interface {
	SendEmail(to string)
	FindEmail(ctx context.Context) (*model.Activity, error)
}
