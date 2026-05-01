package handlers

import (
	"api-gateway/kafka"
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
)

type AskRequest struct {
	Query string `json:"query"`
}

type AskResponse struct {
	Query  string `json:"query"`
	Answer string `json:"answer"`
}

func AskHandler(w http.ResponseWriter, r *http.Request) {
	var req AskRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"invalid request body"}`, http.StatusBadRequest)
		return
	}

	if req.Query == "" {
		http.Error(w, `{"error":"query can not be empty"}`, http.StatusBadRequest)
		return
	}

	// Generate a job ID to track this request through the pipeline.
	// You'll replace this with a proper UUID or DB-generated ID on Day4.
	jobID := fmt.Sprintf("job-%d", rand.Intn(100000))

	// Publish to Kafka. This is non-blocking from the user's perspective -
	// you're handling the work off, not doing it here.
	if err := kafka.PublishQuery(context.Background(), jobID, req.Query); err != nil {
		http.Error(w, `{"error":"failed to queue query"}`, http.StatusInternalServerError)
		return
	}

	resp := AskResponse{
		Query:  req.Query,
		Answer: "stub: RAG pipeline not connected yet",
	}

	w.Header().Set("Content-Type", "application/json")
	//202 Accepted, not 200 OK
	//200 means "done". 202 means "received, processing async"
	//This commnication the async nature of your API to any client
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(resp)
}
