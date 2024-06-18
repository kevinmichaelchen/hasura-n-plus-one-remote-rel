package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Request represents the expected JSON request structure
type Request struct {
	ID int `json:"id"`
}

// Response represents the JSON response structure
type Response struct {
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/", handleRequest)
	log.Println("Server listening on port 8081...")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := Response{
		Message: fmt.Sprintf("Received ID: %d", req.ID),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}