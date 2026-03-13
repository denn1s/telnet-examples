package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "modernc.org/sqlite"
)

type Message struct {
	ID   int    `json:"id"`
	User string `json:"user"`
	Text string `json:"text"`
}

var db *sql.DB

func getMessages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	rows, _ := db.Query("SELECT id, user, text FROM messages")
	defer rows.Close()

	messages := []Message{}
	for rows.Next() {
		var msg Message
		rows.Scan(&msg.ID, &msg.User, &msg.Text)
		messages = append(messages, msg)
	}

	json.NewEncoder(w).Encode(messages)
}

func postMessage(w http.ResponseWriter, r *http.Request) {
	var msg Message
	json.NewDecoder(r.Body).Decode(&msg)
	defer r.Body.Close()

	result, _ := db.Exec("INSERT INTO messages (user, text) VALUES (?, ?)", msg.User, msg.Text)
	id, _ := result.LastInsertId()
	msg.ID = int(id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(msg)
}

func main() {
	var err error
	db, err = sql.Open("sqlite", "file:chat.db")
	if err != nil {
		fmt.Println("Failed to open database:", err)
		return
	}
	defer db.Close()

	db.Exec(`CREATE TABLE IF NOT EXISTS messages (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user TEXT NOT NULL,
		text TEXT NOT NULL
	)`)

	http.HandleFunc("GET /messages", getMessages)
	http.HandleFunc("POST /messages", postMessage)

	fmt.Println("Chat API listening on 0.0.0.0:3000")
	http.ListenAndServe("0.0.0.0:3000", nil)
}
