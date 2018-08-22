package db

import (
	"log"
	"os"
)

var MongoConnStr = os.Getenv("MONGO")

// Gets environment variables and secrets to connect to databases
func init() {
	if MongoConnStr == "" {
		msg := `You need to set an environment variable for MongoDB in the Docker Compose or Docker Swarm file.
			Example:
			MONGO=mongodb://events-db:27017  --inside docker
			MONGO=mongodb://127.0.0.1:27017  --outside of docker
		`
		log.Fatal(msg)
	}
}
