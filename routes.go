package main

import (
	"encoding/json"
	"net/http"
)

func (router *router) setupRoutes() {
	router.HandleFunc("/health", healthCheckHandler)
	router.HandleFunc("/read", readMessageHandler)
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func readMessageHandler(w http.ResponseWriter, r *http.Request) {
	message, _ := QueryMessage()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(message)
}
