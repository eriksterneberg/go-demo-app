package db

import (
	"log"
	"os"
	//"../logging"
	"github.com/eriksterneberg/go-demo-app/events-service/src/logging"
)

var mongoConnStr string


// Gets environment variables and secrets to connect to databases
// Todo:
// For production we need to first check if a secret exists
func init () {
	mongoConnStr = os.Getenv("MONGO")

	if mongoConnStr == "" {
		msg := `You need to set an environment variable for MongoDB in the Docker Compose or Docker Swarm file.
			Example:
			MONGO=mongodb://events-db:27017  --inside docker
			MONGO=mongodb://127.0.0.1:27017  --outside of docker
		`
		log.Fatal(msg)
	}
}

func DatabaseHandlerFactory() DatabaseHandler {
	logging.Debug("Attempting to connect to MongoDB at ", mongoConnStr)

	mongolayer, err := NewMongoDBLayer(mongoConnStr)
	if err != nil {
		log.Fatal("Failed to get mongo db connection")
		os.Exit(1)
	}
	return mongolayer
}
