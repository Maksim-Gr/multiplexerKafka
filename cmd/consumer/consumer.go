package main

import (
	"context"
	"errors"
	"github.com/IBM/sarama"
	"log"
	"strings"
	"sync"
)

type Consumer struct {
	ready chan bool
}

func (consumer *Consumer) Setup(sarama.ConsumerGroupSession) error {
	close(consumer.ready)
	return nil
}

func (consumer *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func main() {
	config := sarama.NewConfig()
	config.ClientID = "multiplexer"
	config.Consumer.Return.Errors = true

	brokers := "localhost:9092"
	topic := "notification"
	group := "test-consumer-group"
	ctx, cancel := context.WithCancel(context.Background())
	client, err := sarama.NewConsumerGroup(strings.Split(brokers, ","), group, config)
	consumer := Consumer{
		ready: make(chan bool),
	}
	if err != nil {
		log.Panicf("Error creating cosnumer group client: %v", err)
	}
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if err := client.Consume(ctx, strings.Split(topic, ","), &consumer); err != nil {
				if errors.Is(err, sarama.ErrClosedConsumerGroup) {
					return
				}
				log.Panicf("error from consumer: %v", err)
			}
			if ctx.Err() != nil {
				return
			}
			consumer.ready = make(chan bool)
		}
	}()
	<-consumer.ready
	log.Println("consumer is up and running")
	cancel()
	wg.Wait()
}

func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message, ok := <-claim.Messages():
			if !ok {
				log.Printf("message channel was closed")
				return nil
			}
			log.Printf(
				"Message claimed: value = %s, timestamp = %v, topic = %s",
				string(message.Value),
				message.Timestamp,
				message.Topic,
			)
			session.MarkMessage(message, "")
		case <-session.Context().Done():
			return nil
		}
	}
}
