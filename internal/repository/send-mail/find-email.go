package sendmail

import (
	"context"
	model "puppet-sync-db/internal/model/local-pelamars"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *Repo) FindEmail(ctx context.Context, limit int64) ([]model.Activity, error) {
	var mailToSend []model.Activity
	filter := bson.D{{Key: "isMail", Value: nil}}
	mongoCollection := r.Db.Collection("activity")

	findOptions := options.Find()
	if limit > 0 {
		findOptions.SetLimit(limit)
	}

	cursor, err := mongoCollection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &mailToSend); err != nil {
		return nil, err
	}

	return mailToSend, nil
}
