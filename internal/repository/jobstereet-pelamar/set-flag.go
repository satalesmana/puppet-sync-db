package jobstereet_pelamar

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *Repo) SetFlag(ctx context.Context, item []primitive.ObjectID) (int64, error) {

	filter := bson.M{"_id": bson.M{"$in": item}}

	mongoCollection := r.Db.Collection(r.cfg.Collection.Remote.Source)
	update := bson.M{
		"$set": bson.M{"sync_status": "1"},
	}

	result, err := mongoCollection.UpdateMany(ctx, filter, update)
	if err != nil {
		return 0, err
	}

	fmt.Printf("Matched %v documents and modified %v documents.\n", result.MatchedCount, result.ModifiedCount)

	return 0, nil
}
