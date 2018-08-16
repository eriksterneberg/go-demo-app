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

func (mongolayer *MongoDBLayer) getSession() *mgo.Session {
	return mongolayer.session.Copy()
}

func (mongolayer MongoDBLayer) AddEvent(e Event) ([]byte, error) {
	session := mongolayer.getSession()
	defer session.Close()

	if !e.ID.Valid() {
		e.ID = bson.NewObjectId()
	}

	if !e.Location.ID.Valid() {
		e.Location.ID = bson.NewObjectId()
	}

	return []byte(e.ID), session.DB(Config.DB).C(Config.EVENTS).Insert(e)
}

func (mongolayer *MongoDBLayer) FindEvent(id []byte) (Event, error) {
	session := mongolayer.getSession()
	defer session.Close()

	event := Event{}
	err := session.DB(Config.DB).C(Config.EVENTS).FindId(bson.ObjectId(id)).One(&event)
	return event, err
}

func (mongolayer *MongoDBLayer) FindEventByName(name string) (Event, error) {
	session := mongolayer.getSession()
	defer session.Close()

	event := Event{}
	err := session.DB(Config.DB).C(Config.EVENTS).Find(bson.M{"name": name}).One(&event)
	return event, err
}

func (mongolayer *MongoDBLayer) FindAllAvailableEvents() ([]Event, error) {
	session := mongolayer.getSession()
	defer session.Close()

	events := []Event{}
	err := session.DB(Config.DB).C(Config.EVENTS).Find(nil).All(&events)
	return events, err
}

func (mongolayer *MongoDBLayer) DeleteEvent(event Event) error {
	session := mongolayer.getSession()
	defer session.Close()

	err := session.DB(Config.DB).C(Config.EVENTS).RemoveId(event.ID)
	return err
}
