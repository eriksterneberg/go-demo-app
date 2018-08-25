// Use factories and interfaces to get publishers and subscribers so it's easier to switch out Kafka

package broker

func MustGetPublisher() EventPublisher {
	broker := &KafkaBroker{}
	broker.Init()
	return broker
}

func MustGetSubscriber() EventSubscriber {
	broker := &KafkaBroker{}
	broker.Init()
	return broker
}
