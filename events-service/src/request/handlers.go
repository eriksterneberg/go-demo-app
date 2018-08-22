package request

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/eriksterneberg/go-demo-app/events-service/src/db"
	"github.com/eriksterneberg/go-demo-app/events-service/src/logging"
	"encoding/hex"
	"github.com/gorilla/mux"
)

// HTTP handler finding a single event
// Tests to add:
// 400 - no SearchCriteria
// 400 - no `search`
// 400 - SearchCriteria not `name` or `id`
// 500 - Internal error should be handled
func FindEventHandler(w http.ResponseWriter, r *http.Request) {
	logging.Debug("FindEventHandler called")

	variables := mux.Vars(r)
	criteria, ok := variables["SearchCriteria"]
	if !ok {
		logging.Debug("User made an invalid call. No SearchCriteria.")
		w.WriteHeader(400)
		fmt.Fprint(w, `{"error": "No valid search criteria found"}`)
		return
	}

	searchkey, ok := variables["search"]
	if !ok {
		logging.Debug("User made an invalid call. No 'search'.")
		w.WriteHeader(400)
		fmt.Fprint(w, `{"error": "No valid search criteria found"}`)
		return
	}

	dbhandler := db.DatabaseHandlerFactory()

	var event db.Event
	var err error
	switch criteria {
	case "name":
		event, err = dbhandler.FindEventByName(searchkey)
	case "id":
		id, err := hex.DecodeString(searchkey)
		if err == nil {
			event, err = dbhandler.FindEvent(id)
		}
	default:
		w.WriteHeader(400)
		fmt.Fprint(w, `{"error": "SearchCriteria must be 'name'' or 'id''."}`)
		return
	}

	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "{\"error\": \"%s\"}", err)
	}

	json.NewEncoder(w).Encode(event)
}


// HTTP handler for getting all events in database
// Tests to add:
// 500 - internal error should be handled
func AllEventHandler(w http.ResponseWriter, r *http.Request) {
	logging.Debug("AllEventHandler called")

	dbhandler := db.DatabaseHandlerFactory()
	events, err := dbhandler.FindAllAvailableEvents()

	if err != nil {
		logging.Error("Encountered error in AllEventHandler:", err)
		w.WriteHeader(500)
		fmt.Fprint(w, `{"errors": "Unknown error occured"}`)
		return
	}

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json;charset=utf8")

	logging.Debug(events)
	json.NewEncoder(w).Encode(&events)
}

// HTTP handler for creating an event
// Tests to add
// 200 - create an event
// 400 - incorrect format
// 500 - internal error should be handled
func CreateEvent(w http.ResponseWriter, r *http.Request) {
	logging.Debug("CreateEvent called")

	event := db.Event{}
	err := json.NewDecoder(r.Body).Decode(&event)

	if err != nil {
		logging.Info("User posted incorrectly formatted event:", event, err)
		w.WriteHeader(400)
		fmt.Fprintf(w, `{"error": "Error occurred while decoding the posted event"}`)
		return
	}

	dbhandler := db.DatabaseHandlerFactory()
	_, err = dbhandler.AddEvent(event)

	if err != nil {
		logging.Error("Got database error while writing event:", event, err)
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"error": "Encountered error while writing to database."}`)
		return
	}

	json.NewEncoder(w).Encode(&event)
}
