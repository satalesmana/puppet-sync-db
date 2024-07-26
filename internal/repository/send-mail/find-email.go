package sendmail

import (
	"context"
	"log"
	model "puppet-sync-db/internal/model/local-pelamars"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *Repo) FindEmail(ctx context.Context, limit int64) (*model.Activity, error) {
	var mailToSend *model.Activity
	filter := bson.D{{Key: "isMail", Value: nil}}
	mongoCollection := r.Db.Collection("activity")

	err := mongoCollection.FindOne(ctx, filter).Decode(&mailToSend)
	if err != nil {
		log.Fatal(err.Error())
	}

	return mailToSend, nil
}
