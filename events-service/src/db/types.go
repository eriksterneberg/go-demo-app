package db

import "gopkg.in/mgo.v2/bson"

type Event struct {
	ID                 bson.ObjectId `bson:"_id"`
	Name               string
	Duration           int
	StartDate, EndDate int64
	Location           Location
}

type Location struct {
	ID                     bson.ObjectId `bson:"_id"`
	Name, Address, Country string
	OpenTime, CloseTime    int
	Halls                  []Hall
}

type Hall struct {
	Name     string `json:"name"`
	Location string `json:"location,omitempty"`
	Capacity int    `json:"capacity"`
}
