package main

import (
	"fmt"
	"net/http"
)

var seriesJSON = `[
  {"id": 1, "name": "Breaking Bad", "seasons": 5},
  {"id": 2, "name": "The Office", "seasons": 9},
  {"id": 3, "name": "Dark", "seasons": 3}
]`

func main() {
	http.HandleFunc("/series", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		fmt.Fprint(w, seriesJSON)
	})

	fmt.Println("JSON API server listening on 0.0.0.0:3000")
	http.ListenAndServe("0.0.0.0:3000", nil)
}
