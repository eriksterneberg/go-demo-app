package request

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eriksterneberg/go-demo-app/party-service/src/db"

	"encoding/json"
	"strings"

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

func mustCreateParty() db.Party {
	dbhandler := db.MustGetDBHandler()

	event := db.Party{
		Name: "Phantom of the Opera",
	}

	id, err := dbhandler.AddParty(event)

	if err != nil {
		log.Panicf("Encountered error initializing test data: %v", err)
	}

	event.ID = bson.ObjectId(id)

	return event
}

func TestGetAllPartysEmpty(t *testing.T) {
	req, _ := http.NewRequest("GET", "/events", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	// No data in db
	if body := response.Body.String(); strings.TrimSpace(body) != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func TestGetAllPartys(t *testing.T) {
	event := mustCreateParty()

	req, _ := http.NewRequest("GET", "/events", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	events := []db.Party{}
	json.NewDecoder(response.Body).Decode(&events)

	dbhandler := db.MustGetDBHandler()
	err := dbhandler.DeleteParty(event)

	if err != nil {
		msg := "Encountered an error when trying to delete from db: %v"
		log.Fatalf(msg, err)
	}

	if length := len(events); length != 1 {
		t.Errorf("GET /events got %d events but expected 1.", length)
	}
}

//func TestGetPartyByName(t *testing.T) {
//	event := mustCreateParty()
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
//func TestGetParty(t *testing.T) {
//
//}
