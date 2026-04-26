package main

import (
	"api-gateway/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/health", handlers.HealthHandler)

	port := ":8080"
	log.Println("Server runing on", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
