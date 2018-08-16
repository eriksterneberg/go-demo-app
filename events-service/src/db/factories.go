package db

import (
	"log"
	"os"
)

func DatabaseHandlerFactory() DatabaseHandler {

	// Todo: put into Docker secret
	mongolayer, err := NewMongoDBLayer("mongodb://127.0.0.1:27017")
	if err != nil {
		log.Fatal("Failed to get mongo db connection")
		os.Exit(1)
	}
	return mongolayer
}
