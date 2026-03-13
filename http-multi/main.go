package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func handleClient(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	// Read the request line (e.g. "GET / HTTP/1.1")
	requestLine, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading request: %v\n", err)
		return
	}

	// Read headers until empty line
	contentLength := 0
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error reading headers: %v\n", err)
			return
		}
		if line == "\r\n" {
			break
		}
		// We need Content-Length to know how much body to read
		if strings.HasPrefix(strings.ToLower(line), "content-length:") {
			fmt.Sscanf(strings.TrimSpace(line[len("content-length:"):]), "%d", &contentLength)
		}
	}

	// Read body if present
	body := ""
	if contentLength > 0 {
		buf := make([]byte, contentLength)
		reader.Read(buf)
		body = string(buf)
	}

	// Parse the request line
	parts := strings.Fields(requestLine)
	if len(parts) < 3 {
		fmt.Printf("Invalid request: %s", requestLine)
		return
	}

	method := parts[0]
	path := parts[1]

	// methdod = 'PUT'
	// path = '/view?series_id=1'
	// location = '/'    query = '?temporary-chat=true' -> queryParams = { temporary-chat: true }

	fmt.Printf("Received %s %s\n", method, path)

	// --- Routing ---
	var response string

	switch {
	case method == "GET" && path == "/":
		response = handleIndex()
	case method == "GET" && path == "/add":
		response = handleAddForm()
	case method == "GET" && path == "/edit":
		response = handleEditForm()
	case method == "POST" && path == "/series":
		response = handleCreate(body)
	case method == "POST" && path == "/series/update":
		response = handleUpdate(body)
	default:
		response = buildResponse("404 Not Found", "<h1>404 - Not Found</h1>")
	}

	conn.Write([]byte(response))
}

// buildResponse is a small helper to format an HTTP response
func buildResponse(status string, html string) string {
	return fmt.Sprintf("HTTP/1.1 %s\r\nContent-Type: text/html\r\nContent-Length: %d\r\n\r\n%s", status, len(html), html)
}

func main() {
	port := ":3000"
	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
		return
	}
	defer listener.Close()

	fmt.Printf("HTTP Server listening on port %s\n", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error accepting connection: %v\n", err)
			continue
		}

		go handleClient(conn)
	}
}
