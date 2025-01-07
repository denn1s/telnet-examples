package main

import (
	"bufio"
	"fmt"
	"net"
)

type Client struct {
	conn net.Conn
	name string
}

func broadcast(message string, clients map[string]Client) {
	for _, client := range clients {
		client.conn.Write([]byte(message))
	}
}

func handleClient(client Client, clients map[string]Client) {
	reader := bufio.NewReader(client.conn)

	// Notify everyone about new client
	broadcast(fmt.Sprintf("* %s joined the chat\n", client.name), clients)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			// Remove client and notify others
			delete(clients, client.name)
			broadcast(fmt.Sprintf("* %s left the chat\n", client.name), clients)
			client.conn.Close()
			return
		}

		// Broadcast the message to all clients
		broadcast(fmt.Sprintf("%s: %s", client.name, message), clients)
	}
}

func main() {
	port := ":2323" // Port with colon prefix
	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
		return
	}
	defer listener.Close()

	clients := make(map[string]Client)
	fmt.Printf("Server listening on port %s\n", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error accepting connection: %v\n", err)
			continue
		}

		// Get client's IP address
		clientIP := conn.RemoteAddr().String()
		client := Client{
			conn: conn,
			name: clientIP,
		}

		// Add to clients map
		clients[clientIP] = client

		// Handle client in a new goroutine
		go handleClient(client, clients)

		fmt.Printf("New connection from %s\n", clientIP)
	}
}
