package local_pelamar

import (
	"context"
	model "puppet-sync-db/internal/model/puppet"

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
	GetSmtpAccount(ctx context.Context) (*model.SmtpAccount, error)
}
