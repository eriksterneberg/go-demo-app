package main

import (
	"net/http"

	"github.com/eriksterneberg/go-demo-app/events-service/src/logging"
	"github.com/eriksterneberg/go-demo-app/events-service/src/request"
)

func main() {
	router := request.GetRouter()

	port := "8080"
	logging.Info("Starting server at port", port)
	logging.Log.Fatal(http.ListenAndServeTLS(":"+port, "cert.pem", "key.pem", router))
}
