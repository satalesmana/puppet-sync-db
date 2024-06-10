package syncdb

import (
	"context"
	"log"
	repoJobstreet "puppet-sync-db/internal/repository/jobstereet-pelamar"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (uc *Uscase) DeleteDBsync() {
	jbPelamarRepo := repoJobstreet.NewRepoHandler(uc.config, uc.mongoConnRemote)
	res, err := jbPelamarRepo.FetchListData(context.Background(), uc.config.Synclimit)
	if err != nil {
		log.Fatalf("error FetchListData", err)
	}

	for _, v := range res {
		if bsonMap, ok := v.(primitive.D); ok {
			var id interface{}
			for _, elem := range bsonMap {
				switch elem.Key {
				case "_id":
					id = elem.Value.(interface{})
				}
			}

			jbPelamarRepo.Delete(context.Background(), id)
		}
	}
}
