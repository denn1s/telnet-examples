package main

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

func main() {
	db, _ := sql.Open("sqlite", "file:test.db")

	db.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER, name TEXT)")
	db.Exec("DELETE FROM users")
	db.Exec("INSERT INTO users VALUES (?, ?)", 1, "Alice")

	var id int
	var name string

	var query = "SELECT id, name FROM users WHERE id = ?"
	db.QueryRow(query, 1).Scan(&id, &name)
	log.Println(id, name)

	db.Exec("INSERT INTO users VALUES (?, ?)", 2, "Bob")
	db.Exec("INSERT INTO users VALUES (?, ?)", 3, "Charlie")
	db.Exec("INSERT INTO users VALUES (?, ?)", 4, "David")

	rows, _ := db.Query("SELECT id, name FROM users WHERE id > ?", 2)

	for rows.Next() {
		var rowID int
		var rowName string
		rows.Scan(&rowID, &rowName)

		log.Println(rowID, rowName)
	}
}
