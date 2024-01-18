package syncdb

import (
	"context"
	"log"
	repoJobstreet "puppet-sync-db/internal/repository/jobstereet-pelamar"
	repoLocalPelamar "puppet-sync-db/internal/repository/local-pelamar"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (uc *Uscase) SyncDBToLocal() {
	jbPelamarRepo := repoJobstreet.NewRepoHandler(uc.mongoConnRemote)

	res, err := jbPelamarRepo.FetchListData(context.Background(), 5)
	if err != nil {
		log.Fatalf("error FetchListData", err)
	}

	var results []interface{}
	for _, v := range res {
		if bsonMap, ok := v.(primitive.D); ok {
			var id interface{}
			for _, elem := range bsonMap {
				switch elem.Key {
				case "_id":
					id = elem.Value.(interface{})
					results = append(results, primitive.D{{Key: "database_id", Value: id}})
				}

			}

		}
	}

	localPelamar := repoLocalPelamar.NewRepoHandler(uc.mongoConnLocal)
	resStore, err := localPelamar.StoreMultiplePelamar(context.Background(), res)
	if err != nil {
		log.Fatalf("error StoreMultiplePelamar", err)
	}

	tes, err := localPelamar.StoreMultipleActivitys(context.Background(), results)
	if err != nil {
		log.Fatalf("error StoreMultiplePelamar", err)
	}

	log.Println("tes", tes)
	log.Println("tes", resStore)
}
