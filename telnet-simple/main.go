package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	port := ":2323"
	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Printf("Server listening on port %s\n", port)

	// Accept a single connection
	conn, err := listener.Accept()
	if err != nil {
		fmt.Printf("Failed to accept connection: %v\n", err)
		os.Exit(1)
	}

	clientAddr := conn.RemoteAddr().String()
	fmt.Printf("Client connected from %s\n", clientAddr)
	defer conn.Close()

	// Start goroutine to read from client
	go func() {
		reader := bufio.NewReader(conn)
		for {
			message, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Client disconnected")
				os.Exit(0)
			}
			fmt.Printf("Received: %s", message)
			// Echo the message back to client
			conn.Write([]byte("Server received: " + message))
		}
	}()

	// Keep main thread alive
	select {}
}
