package db

import (
	"flag"

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

func (m MongoDBLayer) AddEvent(e Event) ([]byte, error) {
	session := m.getSession()
	defer session.Close()

	if !e.ID.Valid() {
		e.ID = bson.NewObjectId()
	}

	if !e.Location.ID.Valid() {
		e.Location.ID = bson.NewObjectId()
	}

	return []byte(e.ID), session.DB(Config.DB).C(Config.EVENTS).Insert(e)
}

func (m *MongoDBLayer) FindEvent(id []byte) (Event, error) {
	session := m.getSession()
	defer session.Close()

	event := Event{}
	err := session.DB(Config.DB).C(Config.EVENTS).FindId(bson.ObjectId(id)).One(&event)
	return event, err
}

func (m *MongoDBLayer) FindEventByName(name string) (Event, error) {
	session := m.getSession()
	defer session.Close()

	event := Event{}
	err := session.DB(Config.DB).C(Config.EVENTS).Find(bson.M{"name": name}).One(&event)
	return event, err
}

func (m *MongoDBLayer) FindAllAvailableEvents() ([]Event, error) {
	session := m.getSession()
	defer session.Close()

	events := []Event{}
	err := session.DB(Config.DB).C(Config.EVENTS).Find(nil).All(&events)
	return events, err
}

func (m *MongoDBLayer) DeleteEvent(e Event) error {
	session := m.getSession()
	defer session.Close()

	err := session.DB(Config.DB).C(Config.EVENTS).RemoveId(e.ID)
	return err
}
