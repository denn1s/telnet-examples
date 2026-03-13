package main

import (
	"fmt"
	"net/url"
)

// handleCreate processes the POST from the "add series" form
func handleCreate(body string) string {
	values, _ := url.ParseQuery(body)

	name := values.Get("name")
	seasons := values.Get("seasons")

	// In a real app we'd save to a database — for now just print it
	fmt.Printf("CREATE series -> name=%q, seasons=%s\n", name, seasons)

	html := fmt.Sprintf(`<html>
<body>
  <h1>Series Created!</h1>
  <p>Name: %s</p>
  <p>Seasons: %s</p>
  <a href="/">Back to list</a>
</body>
</html>`, name, seasons)

	return buildResponse("200 OK", html)
}

// handleUpdate processes the POST (simulating PUT) from the "edit series" form
func handleUpdate(body string) string {
	values, _ := url.ParseQuery(body)

	id := values.Get("id")
	name := values.Get("name")
	seasons := values.Get("seasons")

	// In a real app we'd update the database — for now just print it
	fmt.Printf("UPDATE series -> id=%s, name=%q, seasons=%s\n", id, name, seasons)

	html := fmt.Sprintf(`<html>
<body>
  <h1>Series Updated!</h1>
  <p>ID: %s</p>
  <p>Name: %s</p>
  <p>Seasons: %s</p>
  <a href="/">Back to list</a>
</body>
</html>`, id, name, seasons)

	return buildResponse("200 OK", html)
}
