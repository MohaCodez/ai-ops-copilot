package main

import (
	"api-gateway/handlers"
	"api-gateway/kafka"
	"api-gateway/middleware"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
)

func main() {
	//Initialize Kafka producer once at startup
	// All handlers share this single writer instance.
	kafka.InitProducer("localhost:9092", "queries")
	r := chi.NewRouter()

	r.Use(chiMiddleware.Recoverer)
	r.Use(middleware.Logger)

	r.Get("/health", handlers.HealthHandler)
	r.Post("/ask", handlers.AskHandler)

	log.Println("Server running on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Server failed %v", err)
	}
}
