package main

import (
	"os"

	"fmt"

	broker2 "github.com/eriksterneberg/go-demo-app/event-log/src/broker"
)

func main() {
	broker := broker2.MustGetPublisher()
	defer broker.Close()

	id := os.Args[1]
	fmt.Println("Renaming party with id", id)

	event := broker2.PartyEvent{
		ID:   id,
		Name: "Party of the century",
	}

	partition, offset, err := broker.Publish(event, "RenamedParty")

	fmt.Println("Partition:", partition)
	fmt.Println("Offset:", offset)

	if err != nil {
		panic("Got an error trying to publish to the event log: " + err.Error())
	}
}
