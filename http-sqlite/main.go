package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "modernc.org/sqlite"
)

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("sqlite", "file:series.db")
	if err != nil {
		fmt.Printf("Failed to open database: %v\n", err)
		return
	}

	// Create table and seed data if empty
	db.Exec(`CREATE TABLE IF NOT EXISTS series (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		seasons INTEGER NOT NULL
	)`)

	var count int
	db.QueryRow("SELECT COUNT(*) FROM series").Scan(&count)
	if count == 0 {
		db.Exec("INSERT INTO series (name, seasons) VALUES (?, ?)", "Breaking Bad", 5)
		db.Exec("INSERT INTO series (name, seasons) VALUES (?, ?)", "The Office", 9)
		db.Exec("INSERT INTO series (name, seasons) VALUES (?, ?)", "Dark", 3)
	}

	// Go 1.22+ allows "METHOD /path" patterns
	http.HandleFunc("GET /series", getSeries)
	http.HandleFunc("GET /series/{id}", getOneSerie)
	http.HandleFunc("POST /series", createSerie)
	http.HandleFunc("PUT /series/{id}", updateSerie)
	http.HandleFunc("DELETE /series/{id}", deleteSerie)

	fmt.Println("JSON API server listening on 0.0.0.0:3000")
	http.ListenAndServe("0.0.0.0:3000", nil)
}
