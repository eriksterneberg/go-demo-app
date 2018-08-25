package main

import (
	"fmt"

	broker2 "github.com/eriksterneberg/go-demo-app/event-log/src/broker"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	broker := broker2.MustGetPublisher()
	defer broker.Close()

	id := bson.NewObjectId().Hex()
	fmt.Println("Created party with ID", id)

	event := broker2.PartyEvent{
		ID:   id,
		Name: "ABBA",
	}

	partition, offset, err := broker.Publish(event, "PlannedParty")

	fmt.Println("Partition:", partition)
	fmt.Println("Offset:", offset)

	if err != nil {
		panic("Got an error trying to publish to the event log: " + err.Error())
	}
}
