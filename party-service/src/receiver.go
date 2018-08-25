package main

import (
	"fmt"
	"log"

	"bytes"

	"time"

	"os"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Unable to create channel")
	defer ch.Close()

	queue, err := ch.QueueDeclare(
		"",
		false, // won't be persisted to disk
		false,
		true, // will be deleted upon disconnect
		false,
		nil)

	failOnError(err, "Unable to declare a queue")

	// Need to tell the exchange about this queue
	err = ch.QueueBind(
		queue.Name,
		"event.event.created",
		"events",
		false,
		nil,
	)

	msgs, err := ch.Consume(
		queue.Name,
		"",
		false, // Message is deleted in queue when taken if True
		false,
		false,
		false,
		nil)

	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			fmt.Printf("Got message: %s", d.Body)
			dotCount := bytes.Count(d.Body, []byte("."))
			t := time.Duration(dotCount)
			time.Sleep(t * time.Second)
			log.Printf(" Done\n")
			d.Ack(false) // Remember to always acknowledge message
		}
	}()

	log.Print("Waiting forever for messages on severity level", os.Getenv("SEVERITY"))
	<-forever
}
