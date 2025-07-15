package infrastructure

import (
	"encoding/json"
	"net/http"
)

type ProblemResponse struct {
	Type   string `json:"type"`
	Title  string `json:"title"`
	Status int    `json:"status"`
	Detail string `json:"detail"`
}

func Write(w http.ResponseWriter, p ProblemResponse) {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(p.Status)
	json.NewEncoder(w).Encode(p)
}
