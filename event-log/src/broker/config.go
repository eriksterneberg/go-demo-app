// Defines settings for Kafka, read from the environment

package broker

import (
	"os"
	"strings"
)

var kafkaBrokerListString = os.Getenv("KAFKA_BROKERS")

var KafkaBrokerList []string

func init() {
	if kafkaBrokerListString == "" {
		KafkaBrokerList = []string{"0.0.0.0:9092"}
	} else {
		KafkaBrokerList = strings.Split(kafkaBrokerListString, ",")
	}
}
