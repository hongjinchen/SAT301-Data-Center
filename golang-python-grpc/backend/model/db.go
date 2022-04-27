package model

/*
The DB object.
*/

import (
	"context"
	"log"
	"sync"
	"time"

	"aruix.net/chj-fyp/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	// MongoURL string = "mongodb://chj:hongchenjin@mongo:27017/?authSource=admin&readPreference=primary&ssl=false"
	MongoURL string = config.Config.DB.MongoURL
	DATABASE string = config.Config.DB.DBName
)

var (
	Wg sync.WaitGroup

	Database *mongo.Database = DataBase()
)

func DataBase() *mongo.Database {
	Wg.Add(1)
	// Wait for the database get ready, the router (controller) can start.
	defer Wg.Done()
	client, err := mongo.NewClient(options.Client().ApplyURI(MongoURL))

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB")
	return client.Database(DATABASE)
}

func Collection(coll string) *mongo.Collection {
	collection := Database.Collection(coll)
	return collection
}
