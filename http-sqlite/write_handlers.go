package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// POST /series — create a new serie
func createSerie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var s Serie
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{"error": "invalid JSON"}`)
		return
	}

	result, err := db.Exec("INSERT INTO series (name, seasons) VALUES (?, ?)", s.Name, s.Seasons)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "%s"}`, err)
		return
	}

	id, _ := result.LastInsertId()
	s.ID = int(id)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(s)
}

// PUT /series/{id} — update an existing serie
func updateSerie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	id := r.PathValue("id")

	var s Serie
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{"error": "invalid JSON"}`)
		return
	}

	result, err := db.Exec("UPDATE series SET name = ?, seasons = ? WHERE id = ?", s.Name, s.Seasons, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "%s"}`, err)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, `{"error": "serie not found"}`)
		return
	}

	fmt.Sscanf(id, "%d", &s.ID)
	json.NewEncoder(w).Encode(s)
}
