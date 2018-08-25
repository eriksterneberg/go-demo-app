package request

import "github.com/gorilla/mux"

func GetRouter() *mux.Router {
	router := mux.NewRouter()
	eventsrouter := router.PathPrefix("/events").Subrouter()
	eventsrouter.Methods("GET").Path("/{SearchCriteria}/{search").HandlerFunc(FindPartyHandler)
	eventsrouter.Methods("GET").Path("").HandlerFunc(AllPartyHandler)
	return router
}
