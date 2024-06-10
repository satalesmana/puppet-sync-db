package jobstereet_pelamar

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *Repo) Delete(ctx context.Context, id interface{}) error {
	filter := bson.D{{Key: "_id", Value: id}}
	mongoCollection := r.Db.Collection(r.cfg.Collection.Remote.Source)

	_, err := mongoCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}
