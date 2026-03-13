package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Message struct {
	User string `json:"user"`
	Text string `json:"text"`
}

var messages []Message

func getMessages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}

func postMessage(w http.ResponseWriter, r *http.Request) {
	var msg Message
	json.NewDecoder(r.Body).Decode(&msg)
	defer r.Body.Close()

	messages = append(messages, msg)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(msg)
}

func main() {
	messages = []Message{}

	http.HandleFunc("GET /messages", getMessages)
	http.HandleFunc("POST /messages", postMessage)

	fmt.Println("Chat API listening on 0.0.0.0:3000")
	http.ListenAndServe("0.0.0.0:3000", nil)
}
