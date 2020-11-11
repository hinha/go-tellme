package platform

import (
	"context"
	"github.com/Kamva/mgm/v3"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectionDatabase(uri, dbName, domain string) *mongo.Database {
	logrus.WithFields(logrus.Fields{
		"platform": "mongo",
		"domain":   domain,
	}).Info("Connection mongodb")

	opClient := options.Client().ApplyURI(uri)
	err := mgm.SetDefaultConfig(nil, dbName, opClient)
	if err != nil {
		logrus.Error(err)
	}

	_, client, db, _ := mgm.DefaultConfigs()
	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		logrus.Error(err)
	}

	logrus.Info("Connected to MongoDB!")
	return db
}
