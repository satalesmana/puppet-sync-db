package main

import (
	"log"
	"puppet-sync-db/internal/config"
)

func main() {
	ConfigSyncDb, err := config.New("puppet-sync-db")
	if err != nil {
		panic(err)
	}

	log.Fatal(startApp(ConfigSyncDb))
}
