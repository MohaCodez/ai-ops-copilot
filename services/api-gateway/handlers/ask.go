package handlers

import (
	"encoding/json"
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

	resp := AskResponse{
		Query:  req.Query,
		Answer: "stub: RAG pipeline not connected yet",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
