// Simple interfaces for components that pass messages between services through a message broker.

package broker

type EventInterface interface {
	TopicName() string
}

type Broker interface {
	Init()
	Close()
}

type EventPublisher interface {
	Broker
	Publish(events EventInterface, key string) (partition int32, offset int64, err error)
}

type EventSubscriber interface {
	Broker
	Subscribe(topic string, output chan<- Message, offset int64) error
}
