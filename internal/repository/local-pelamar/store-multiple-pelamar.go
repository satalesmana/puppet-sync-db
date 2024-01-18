package local_pelamar

import (
	"context"
	"log"
)

func (r *Repo) StoreMultiplePelamar(ctx context.Context, data []interface{}) (int64, error) {
	mongoCollection := r.Db.Collection("database")
	_, err := mongoCollection.InsertMany(ctx, data)
	if err != nil {
		log.Fatal(err)
	}

	return 0, nil
}

func (r *Repo) StoreMultipleActivitys(ctx context.Context, data []interface{}) (int64, error) {
	mongoCollection := r.Db.Collection("activity")
	_, err := mongoCollection.InsertMany(ctx, data)
	if err != nil {
		log.Fatal(err)
	}

	return 0, nil
}
