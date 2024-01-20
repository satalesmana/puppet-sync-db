package jobstereet_pelamar

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *Repo) FetchListData(ctx context.Context, limit int64) ([]interface{}, error) {
	filter := bson.D{{Key: "sync_status", Value: nil}}
	mongoCollection := r.Db.Collection(r.cfg.Collection.Remote.PrimaryCollection)
	findOptions := options.Find()
	findOptions.SetLimit(limit)
	cursor, err := mongoCollection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []interface{}
	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	return results, nil
}
