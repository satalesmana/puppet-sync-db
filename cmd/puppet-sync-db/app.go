package main

import (
	"log"
	config "puppet-sync-db/internal/config"
	"puppet-sync-db/internal/usecase/syncdb"
)

func startApp(config *config.Config) error {

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

	syncDB := syncdb.NewSyncDBHandler(mongoConnRemote, mongoConnLocal)
	syncDB.SyncDBToLocal()

	return nil
}
