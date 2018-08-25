package db

import (
	"log"

	"github.com/eriksterneberg/go-demo-app/party-service/src/logging"
)

// Abstracts the database layer so it's easier to switch out one data storage for another
func MustGetDBHandler() DatabaseHandler {
	logging.Debug("Attempting to connect to MongoDB at ", MongoConnStr)

	dbhandler, err := NewMongoDBLayer(MongoConnStr)

	if err != nil {
		log.Fatalf("Failed to get mongo db connection: %s", err)
	}
	return dbhandler
}
