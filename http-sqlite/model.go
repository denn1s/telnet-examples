package main

// Serie represents a TV series — this struct maps to both
// the database columns and the JSON fields
type Serie struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Seasons int    `json:"seasons"`
}
