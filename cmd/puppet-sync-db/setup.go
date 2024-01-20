package main

import (
	"context"
	"fmt"
	"log"
	model "puppet-sync-db/internal/model/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func setupMongoDBConnection(config *model.Config, dbType string) (*mongo.Database, error) {

	var (
		mongodbURI  string
		mongoConfig = config.MongoDBRemote
	)

	if dbType == "local" {
		mongoConfig = config.MongoDBLocal
	}

	if mongoConfig.Username == "" || mongoConfig.Password == "" {
		mongodbURI = fmt.Sprintf("mongodb://%s:%s", mongoConfig.Host, mongoConfig.Port)
	} else {
		mongodbURI = fmt.Sprintf("mongodb+srv://%s:%s@%s", mongoConfig.Username, mongoConfig.Password, mongoConfig.Host)
	}

	fmt.Println("mongodbURI", mongodbURI)
	clientOptions := options.Client().ApplyURI(mongodbURI)
	mongoClient, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	database := mongoClient.Database(mongoConfig.Database)

	err = mongoClient.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully connected to MongoDB")
	return database, nil
}
