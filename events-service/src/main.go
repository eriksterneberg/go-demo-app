package main

import (
	"net/http"
	//"./logging"
	"github.com/eriksterneberg/go-demo-app/events-service/src/logging"

	//"./request"
	"github.com/eriksterneberg/go-demo-app/events-service/src/request"

	"github.com/gorilla/mux"
)

func GetRouter() *mux.Router {
	router := mux.NewRouter()
	eventsrouter := router.PathPrefix("/events").Subrouter()
	eventsrouter.Methods("GET").Path("/{SearchCriteria}/{search").HandlerFunc(request.FindEventHandler)
	eventsrouter.Methods("GET").Path("").HandlerFunc(request.AllEventHandler)
	eventsrouter.Methods("POST").Path("").HandlerFunc(request.CreateEvent)
	return router
}

func main() {
	// Init router
	router := GetRouter()

	port := "8080"
	logging.Info("Starting server at port", port)
	logging.Log.Fatal(http.ListenAndServe(":"+port, router))
}
