package main

import (
	"fmt"
	"io"
	"net/http"
)

var chatAPI = "http://localhost:3000"

func getMessages(w http.ResponseWriter, r *http.Request) {
	resp, _ := http.Get(chatAPI + "/messages")
	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	io.Copy(w, resp.Body)
}

func postMessage(w http.ResponseWriter, r *http.Request) {
	resp, _ := http.Post(chatAPI+"/messages", "application/json", r.Body)
	defer resp.Body.Close()
	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

func main() {
	http.Handle("GET /", http.FileServer(http.Dir("static")))
	http.HandleFunc("GET /api/messages", getMessages)
	http.HandleFunc("POST /api/messages", postMessage)

	fmt.Println("Chat client listening on 0.0.0.0:8000")
	http.ListenAndServe("0.0.0.0:8000", nil)
}
