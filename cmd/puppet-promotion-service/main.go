package main

import (
	"log"
	"puppet-sync-db/internal/config"
)

func main() {
	ConfigSyncDb, err := config.New("puppet-promotion-service")
	if err != nil {
		panic(err)
	}

	log.Fatal(startApp(ConfigSyncDb))
}
