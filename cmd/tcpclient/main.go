package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ravidhavlesha/go-multiproto-kvs/pkg/protocol/tcp"
)

func main() {
	// Create a new TCP client
	tcpClient := tcp.NewTCPClient(":8080")

	// Establish connection to the TCP server
	if err := tcpClient.Connect(); err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer tcpClient.Close()

	// Create a reader for standard input (for user commands)
	reader := bufio.NewReader(os.Stdin)

	// Command loop
	for {
		fmt.Print("> ")
		// Read user input
		userInput, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Failed to read user input: %v", err)
			continue
		}

		userInput = strings.TrimSpace(userInput)

		// If the user types "exit", break the loop and close the connection
		if strings.ToLower(userInput) == "exit" {
			fmt.Println("Exiting client.")
			break
		}

		// Send commands to the server
		if err := tcpClient.Send(userInput); err != nil {
			log.Printf("Failed to send command: %v", err)
			continue
		}

		// Wait for the server's response and print it
		response, err := tcpClient.Receive()
		if err != nil {
			log.Printf("Failed to receive response: %v", err)
			continue
		}
		fmt.Printf("Server response: %s\n", response)
	}
}
