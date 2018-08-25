// Kafka implementation of Broker, Publisher and Subscriber interfaces

package broker

import (
	"encoding/json"

	"github.com/Shopify/sarama"
)

type KafkaBroker struct {
	Client sarama.Client
}

func (r *KafkaBroker) Init() {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	client, err := sarama.NewClient(KafkaBrokerList, config)

	if err != nil {
		panic("Unable to instantiate Sarama Kafka client: " + err.Error())
	}

	r.Client = client
}

func (k *KafkaBroker) Close() {
	k.Client.Close()
}

// Takes a single EventInterface event and publishes it. The topic is derived from the event.
func (k *KafkaBroker) Publish(event EventInterface, key string) (partition int32, offset int64, err error) {
	producer, err := sarama.NewSyncProducerFromClient(k.Client)
	if err != nil {
		return 0, 0, err
	}

	jsonBody, err := json.Marshal(event)
	if err != nil {
		return 0, 0, err
	}

	// On Partition
	// Should be set to different values for different data, to balance load.
	// For instance, you can calculate a hash from a user ID. That will guarantee that the events
	// for that user will be replayed in the correct order as events in a partition are ordered.
	msg := &sarama.ProducerMessage{
		Topic:     event.TopicName(),
		Key:       sarama.StringEncoder(key),
		Value:     sarama.ByteEncoder(jsonBody),
		Partition: 0,
	}

	partition, offset, err = producer.SendMessage(msg)

	return partition, offset, err
}

// Subscribes to a certain topic and outputs results to a channel.
// Client needs to remember the offset itself. Kafka only stores the offset of the latest commit,
// but if the client was offline for a while it could have missed messages and as a consequence
// needs to track the offset for itself.
func (k *KafkaBroker) Subscribe(topic string, output chan<- Message, offset int64) (err error) {
	consumer, err := sarama.NewConsumerFromClient(k.Client)

	if err != nil {
		panic("Got an error while trying to create a consumer: " + err.Error())
	}

	conn, err := consumer.ConsumePartition(
		topic,
		0,
		offset, // Start from the next unread message
	)

	if err != nil {
		panic("Got an error while trying to consume a partition: " + err.Error())
	}

	go func() {
		for msg := range conn.Messages() {
			output <- Message{
				Key:    msg.Key,
				Value:  msg.Value,
				Offset: msg.Offset,
			}
		}
	}()

	return err
}
