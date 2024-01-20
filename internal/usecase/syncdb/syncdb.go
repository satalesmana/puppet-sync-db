package syncdb

import (
	"context"
	"log"
	repoJobstreet "puppet-sync-db/internal/repository/jobstereet-pelamar"
	repoLocalPelamar "puppet-sync-db/internal/repository/local-pelamar"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (uc *Uscase) SyncDBToLocal() {
	//fetch data from remote server
	jbPelamarRepo := repoJobstreet.NewRepoHandler(uc.config, uc.mongoConnRemote)
	res, err := jbPelamarRepo.FetchListData(context.Background(), 5)
	if err != nil {
		log.Fatalf("error FetchListData", err)
	}

	//store data to local server
	localPelamar := repoLocalPelamar.NewRepoHandler(uc.mongoConnLocal)
	resStore, err := localPelamar.StoreMultiplePelamar(context.Background(), res)
	if err != nil {
		log.Fatalf("error StoreMultiplePelamar", err)
	}

	//prepare / parsing clean data
	var results []interface{}
	for _, v := range res {
		if bsonMap, ok := v.(primitive.D); ok {
			var id interface{}
			var name string
			var email string
			var phone string

			for _, elem := range bsonMap {
				switch elem.Key {
				case "_id":
					id = elem.Value.(interface{})
				case "phone":
					phone = strings.Replace(elem.Value.(string), " ", "", 1)
				case "email":
					email = elem.Value.(string)
				case "firstName":
					name = elem.Value.(string)
				case "lastName":
					name += " " + elem.Value.(string)
				}

			}

			results = append(results, primitive.D{
				{Key: "database_id", Value: id},
				{Key: "name", Value: name},
				{Key: "email", Value: email},
				{Key: "phone", Value: phone},
				{Key: "isMail", Value: nil},
				{Key: "isPhone", Value: nil},
			})

		}
	}

	//store clean data to local server
	tes, err := localPelamar.StoreMultipleActivitys(context.Background(), results)
	if err != nil {
		log.Fatalf("error StoreMultiplePelamar", err)
	}

	log.Println("tes", tes)
	log.Println("tes", resStore)
}
