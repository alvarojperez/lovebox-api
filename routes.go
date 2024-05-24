package main

import (
	"encoding/json"
	"net/http"
)

func (router *router) setupRoutes() {
	router.HandleFunc("/read", readMessageHandler)
}

func readMessageHandler(w http.ResponseWriter, r *http.Request) {
	message, _ := QueryMessage()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(message)
}
