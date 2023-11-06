package main

import (
	"github.com/IBM/sarama"
	"strings"
)

// Temporary vars for testing
const (
	ConsumerGroup      = "multiplexer-group"
	ConsumerTopic      = "notifications"
	ConsumerPort       = ":8081"
	KafkaServerAddress = "localhost:9092"
)

func main() {
	config := sarama.NewConfig()
	config.ClientID = "multiplexer"
	config.Consumer.Return.Errors = true

	brokers := "localhost:9092"
	consumer, err := sarama.NewConsumerGroup(strings.Split(brokers, ","), ConsumerGroup, config)
}

type Consumer struct {
	ready chan bool
}

func (consumer *Consumer) ConsumeClaim() {

}
