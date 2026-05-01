package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
)

// QueryMessage mirrors exactly what the api-gateway puublishes.
// Both sides must agree on this shape - this is your internal message contract
type QueryMessage struct {
	JobID string `json:"job_id"`
	Query string `json:"query"`
}

func main() {
	// NewReader creates a Kafka consumer.
	// GroupID is critical - it tells Kafka to track which messages
	// this consumer group has already processed.
	// If you run two worker instances, Kafka splits messages between them.
	// Without GroupID, every worker would process every message.
	reader := kafka.NewReader((kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "queries",
		GroupID: "worker-group",
	}))
	defer reader.Close()

	log.Println("Worker started - listening on topic: queries")

	for {
		// ReadMessage blocks here until a message arrives.
		// This is efficient - the worker sleeps until there's work to do
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Error reading message: %v", err)
			continue // log and keep going, never crash the loop
		}

		var query QueryMessage
		if err := json.Unmarshal(msg.Value, &query); err != nil {
			log.Printf("Error parsing message: %v", err)
			continue
		}

		// Stub processing for now.
		// Day 4: this is where you'll call RAG service via gRPC
		// and store the result in Postgres
		log.Printf("Processing job - ID: %s, Query: %s", query.JobID, query.Query)
	}
}
