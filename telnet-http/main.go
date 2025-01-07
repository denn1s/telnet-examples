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

	// Read the HTTP request
	requestLine, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading request: %v\n", err)
		return
	}

	// Read headers until we get an empty line (denoted by \r\n)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error reading headers: %v\n", err)
			return
		}

		// Empty line (just \r\n) marks the end of headers
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
	protocol := parts[2]

	fmt.Printf("Received %s request for %s using %s\n", method, path, protocol)

	// Prepare HTTP response
	response := "HTTP/1.1 200 OK\r\n" +
		"Content-Type: text/plain\r\n" +
		"\r\n" +
		"Hello World!\r\n" +
		"Path: " + path

	// Send response
	conn.Write([]byte(response))
}

func main() {
	port := ":8080"
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

		clientIP := conn.RemoteAddr().String()
		fmt.Printf("New connection from %s\n", clientIP)

		// Handle client in a new goroutine
		go handleClient(conn)
	}
}
