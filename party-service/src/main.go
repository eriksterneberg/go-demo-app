package main

import (
	"net/http"

	"os"

	"encoding/json"

	"fmt"

	"io/ioutil"

	"strings"

	"strconv"

	broker2 "github.com/eriksterneberg/go-demo-app/event-log/src/broker"
	"github.com/eriksterneberg/go-demo-app/party-service/src/logging"
	"github.com/eriksterneberg/go-demo-app/party-service/src/request"
)

// Subscribes to interesting events in the event log and persists to local storage
func subscribe() {
	// Read previous offset if it exists
	var offset int
	var offset64 int64
	raw, err := ioutil.ReadFile("party.offset")

	if err != nil {
		offset = 0
		fmt.Println("Offset not on disk. Setting to zero.")
	} else {
		offset, _ = strconv.Atoi(strings.TrimSpace(string(raw)))
		offset64 = int64(offset) + 1
		fmt.Println("Read offset from disk:", offset64)
	}

	// Spin off goroutine that listens for events and saves to storage
	broker := broker2.MustGetSubscriber()
	defer broker.Close()

	// Subscribe to all CUD operations related to events
	output := make(chan broker2.Message)

	go func() {
		err := broker.Subscribe(broker2.PartyEvent{}.TopicName(), output, offset64)
		if err != nil {
			logging.Log.Fatalf("Encountered error while subscribing to events: " + err.Error())
		}
	}()

	var lastOffset int64

	// Listen for events and persist to database
	for msg := range output {
		logging.Infof("Received a %s message from the party topic.", msg.Key)
		lastOffset = msg.Offset

		event := broker2.PartyEvent{}
		err := json.Unmarshal(msg.Value, &event)

		if err != nil {
			logging.Error("Got error while unmarshaling PartyEvent message:", msg)
			continue
		}
		//party := db.Party{
		//	ID:   bson.ObjectId(event.ID),
		//	Name: event.Name,
		//}

		//party.Ingest(event.OperationID, msg.Offset)

		//err = updateOffset(msg.Offset)
		//if err != nil {
		//	logging.Error("Got an error while updating offset to disk:", err)
		//}

		if lastOffset%10 == 0 { // Magic number
			fmt.Printf("Persist last offset to disk: %d, %[1]T\n", lastOffset)
			ioutil.WriteFile(
				"party.offset", // make constant
				[]byte(fmt.Sprint(lastOffset)),
				0644,
			)
		}
	}

}

func main() {
	logging.Info("Starting to listen for events on the message broker")
	go subscribe()

	router := request.GetRouter()
	var port = os.Getenv("PORT")
	defer fmt.Println("test, test")

	logging.Info("Starting server at port", port)
	logging.Log.Fatal(http.ListenAndServeTLS(":"+port, "cert.pem", "key.pem", router))
}
