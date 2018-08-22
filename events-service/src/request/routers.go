package request

import "github.com/gorilla/mux"

func GetRouter() *mux.Router {
	router := mux.NewRouter()
	eventsrouter := router.PathPrefix("/events").Subrouter()
	eventsrouter.Methods("GET").Path("/{SearchCriteria}/{search").HandlerFunc(FindEventHandler)
	eventsrouter.Methods("GET").Path("").HandlerFunc(AllEventHandler)
	eventsrouter.Methods("POST").Path("").HandlerFunc(CreateEvent)
	return router
}