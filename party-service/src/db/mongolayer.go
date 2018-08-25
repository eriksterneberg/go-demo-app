package db

import (
	"flag"

	"github.com/eriksterneberg/go-demo-app/party-service/src/logging"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var Config struct {
	DB, USERS, EVENTS string
}

// Rename database for testing so no one runs UT/FT on production database by mistake
func init() {
	if flag.Lookup("test.v") == nil {
		Config.DB = "eventsDB"
		Config.EVENTS = "events"
	} else {
		Config.DB = "eventsDB_test"
		Config.EVENTS = "events_test"
	}
}

type MongoDBLayer struct {
	session *mgo.Session
}

func NewMongoDBLayer(connection string) (*MongoDBLayer, error) {
	s, err := mgo.Dial(connection)
	if err != nil {
		return nil, err
	}
	return &MongoDBLayer{session: s}, err
}

func (m *MongoDBLayer) getSession() *mgo.Session {
	return m.session.Copy()
}

func (m MongoDBLayer) AddParty(e Party) ([]byte, error) {
	session := m.getSession()
	defer session.Close()

	if !e.ID.Valid() {
		e.ID = bson.NewObjectId()
	}

	//if !e.Location.ID.Valid() {
	//	e.Location.ID = bson.NewObjectId()
	//}

	return []byte(e.ID), session.DB(Config.DB).C(Config.EVENTS).Insert(e)
}

func (m *MongoDBLayer) FindParty(id []byte) (Party, error) {
	session := m.getSession()
	defer session.Close()

	event := Party{}
	err := session.DB(Config.DB).C(Config.EVENTS).FindId(bson.ObjectId(id)).One(&event)
	return event, err
}

func (m *MongoDBLayer) FindPartyByName(name string) (Party, error) {
	session := m.getSession()
	defer session.Close()

	event := Party{}
	err := session.DB(Config.DB).C(Config.EVENTS).Find(bson.M{"name": name}).One(&event)
	return event, err
}

func (m *MongoDBLayer) FindAllAvailablePartys() ([]Party, error) {
	session := m.getSession()
	defer session.Close()

	events := []Party{}
	err := session.DB(Config.DB).C(Config.EVENTS).Find(nil).All(&events)
	return events, err
}

func (m *MongoDBLayer) DeleteParty(e Party) error {
	session := m.getSession()
	defer session.Close()

	err := session.DB(Config.DB).C(Config.EVENTS).RemoveId(e.ID)
	return err
}

func (e Party) Ingest(operationID int) error {
	logging.Debug("Preparing to ingest Create, Update, Delete message for party events")

	// First we need to save the latest offset id somewhere and don't download more messages than we need

	return nil
}
