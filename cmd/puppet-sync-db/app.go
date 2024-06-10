package main

import (
	"log"
	model "puppet-sync-db/internal/model/config"
	"puppet-sync-db/internal/usecase/syncdb"
)

func startApp(config *model.Config) error {

	mongoConnRemote, err := setupMongoDBConnection(config, "remote")
	if err != nil {
		log.Fatalf("error startup service when connection setup to mongodb. err: %v", err)
		return err
	}

	mongoConnLocal, err := setupMongoDBConnection(config, "local")
	if err != nil {
		log.Fatalf("error startup service when connection setup to mongodb. err: %v", err)
		return err
	}

	syncDB := syncdb.NewSyncDBHandler(config, mongoConnRemote, mongoConnLocal)
	// syncDB.SyncDBToLocal()
	syncDB.DeleteDBsync()

	return nil
}
