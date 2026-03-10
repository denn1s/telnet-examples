package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

var seriesJSON = `[
  {"id": 1, "name": "Breaking Bad", "seasons": 5},
  {"id": 2, "name": "The Office", "seasons": 9},
  {"id": 3, "name": "Dark", "seasons": 3}
]`

func handleClient(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	// Read the request line (e.g. "GET /series HTTP/1.1")
	requestLine, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading request: %v\n", err)
		return
	}

	// Read headers until empty line
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error reading headers: %v\n", err)
			return
		}
		if line == "\r\n" {
			break
		}
	}

	// Parse the request line
	parts := strings.Fields(requestLine)
	if len(parts) < 3 {
		fmt.Printf("Invalid request: %s", requestLine)
		return
	}

	method := parts[0]
	path := parts[1]

	fmt.Printf("Received %s %s\n", method, path)

	// --- Routing ---
	var response string

	switch {
	case method == "GET" && path == "/series":
		response = buildResponse("200 OK", seriesJSON)
	default:
		response = buildResponse("404 Not Found", `{"error": "not found"}`)
	}

	conn.Write([]byte(response))
}

// buildResponse formats an HTTP response with JSON content type
func buildResponse(status string, body string) string {
	return fmt.Sprintf(
		"HTTP/1.1 %s\r\nContent-Type: application/json\r\nAccess-Control-Allow-Origin: *\r\nContent-Length: %d\r\n\r\n%s",
		status, len(body), body,
	)
}

func main() {
	port := "0.0.0.0:3000"
	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
		return
	}
	defer listener.Close()

	fmt.Printf("JSON API server listening on %s\n", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error accepting connection: %v\n", err)
			continue
		}

		go handleClient(conn)
	}
}
