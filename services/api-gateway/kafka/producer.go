package kafka

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

// Writer is the Kafka producer connection
// Package-level so it's initialized once at startup and reused across requests.
// Creating a new writer per request would be wasteful and slow
var Writer *kafka.Writer

// InitProducer sets up the Kafka writer. Call this once in main.go at startup.
// brokerAddress: where Kafka is running (e.x. "localhost:9092")
// topic: which topic t publish to (e.g. "queries")
func InitProducer(brokerAddress string, topic string) {
	Writer = &kafka.Writer{
		Addr:         kafka.TCP(brokerAddress),
		Topic:        topic,
		Balancer:     &kafka.LeastBytes{},
		WriteTimeout: 10 * time.Second,
	}
	log.Printf("Kafka producer initialized - broker: %s, topic: %s", brokerAddress, topic)
}

// PublishQuery sends a query message to Kafka.
// jobId becomes the message key - Kafka uses this for partition routing.
// All messages with the same key go to the same partition, preserving order.
func PublishQuery(ctx context.Context, jobID string, query string) error {
	payload := map[string]string{
		"job_id": jobID,
		"query":  query,
	}

	//Serialize to JSON - this is the raw bytes Kafka stores
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	err = Writer.WriteMessages(ctx, kafka.Message{
		Key:   []byte(jobID),
		Value: data,
	})

	if err != nil {
		return err
	}

	log.Printf("Published to Kafka - jobId %s, query: %s", jobID, query)
	return nil
}
