package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GET /series — list all series
func getSeries(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	rows, err := db.Query("SELECT id, name, seasons FROM series")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "%s"}`, err)
		return
	}
	defer rows.Close()

	series := []Serie{}
	for rows.Next() {
		var s Serie
		rows.Scan(&s.ID, &s.Name, &s.Seasons)
		series = append(series, s)
	}

	json.NewEncoder(w).Encode(series)
}

// GET /series/{id} — get one serie
func getOneSerie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	id := r.PathValue("id")

	var s Serie
	err := db.QueryRow("SELECT id, name, seasons FROM series WHERE id = ?", id).Scan(&s.ID, &s.Name, &s.Seasons)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, `{"error": "serie not found"}`)
		return
	}

	json.NewEncoder(w).Encode(s)
}
