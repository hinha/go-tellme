package platform

import (
	"context"
	"fmt"
	"github.com/Kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func ConnectionDatabase(uri, dbName string) *mongo.Database {
	opClient := options.Client().ApplyURI(uri)
	err := mgm.SetDefaultConfig(nil, dbName, opClient)
	if err != nil {
		log.Fatal(err)
	}

	_, client, db, _ := mgm.DefaultConfigs()
	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	return db
}
