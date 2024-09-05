package local_pelamar

import (
	"context"
	model "puppet-sync-db/internal/model/puppet"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *Repo) GetSmtpAccount(ctx context.Context) (*model.SmtpAccount, error) {
	var result *model.SmtpAccount

	filter := bson.D{{Key: "status", Value: "1"}}
	smtpAccount := r.Db.Collection("smtp_account")

	err := smtpAccount.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
