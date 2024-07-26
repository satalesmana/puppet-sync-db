package sendmail

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *Repo) SetFlagMail(ctx context.Context, email string) (int64, error) {
	filter := bson.M{"email": email}

	mongoCollection := r.Db.Collection("activity")
	update := bson.M{
		"$set": bson.M{"isMail": r.Cfg.Email.SetFlag},
	}

	result, err := mongoCollection.UpdateMany(ctx, filter, update)
	if err != nil {
		return 0, err
	}

	fmt.Printf("Matched %v documents and modified %v documents.\n", result.MatchedCount, result.ModifiedCount)

	return 0, nil
}
