package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"github.com/eriksterneberg/go-demo-app/events-service/src/db"
	"strings"
	"testing"

	"gopkg.in/mgo.v2/bson"
)

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d was different from actual %d.", expected, actual)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	GetRouter().ServeHTTP(rr, req)

	return rr
}

func mustCreateEvent() db.Event {
	dbhandler := db.DatabaseHandlerFactory()

	event := db.Event{
		Name:      "Phantom of the Opera",
		Duration:  7200000,
		StartDate: 123,
		EndDate:   124,
	}

	id, err := dbhandler.AddEvent(event)

	if err != nil {
		log.Panicf("Encountered error initializing test data: %v", err)
	}

	// Clean up
	defer dbhandler.DeleteEvent(id)

	event.ID = bson.ObjectId(id)

	return event
}

func TestGetAllEventsEmpty(t *testing.T) {
	req, _ := http.NewRequest("GET", "/events", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	// No data in db
	if body := response.Body.String(); strings.TrimSpace(body) != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func TestGetAllEvents(t *testing.T) {
	mustCreateEvent()

	req, _ := http.NewRequest("GET", "/events", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	events := []db.Event{}
	json.NewDecoder(response.Body).Decode(&events)
	fmt.Println(events)

	if length := len(events); length != 1 {
		t.Errorf("GET /events got %d events but expected 1.", length)
	}

}

//func TestGetEventByName(t *testing.T) {
//	event := mustCreateEvent()
//
//	url := fmt.Sprintf("/events/name/%s", event.Name)
//	req, _ := http.NewRequest("GET", url, nil)
//	response := executeRequest(req)
//	checkResponseCode(t, http.StatusOK, response.Code)
//
//	events := []db.
//
//}
//
//func TestGetEvent(t *testing.T) {
//
//}
