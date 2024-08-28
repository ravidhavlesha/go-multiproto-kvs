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
	tcpClient := tcp.NewTCPClient(":8080")

	if err := tcpClient.Connect(); err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer tcpClient.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		userInput, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Failed to read user input: %v", err)
			continue
		}

		userInput = strings.TrimSpace(userInput)

		if strings.ToLower(userInput) == "exit" {
			fmt.Println("Exiting client.")
			break
		}

		if err := tcpClient.Send(userInput); err != nil {
			log.Printf("Failed to send command: %v", err)
			continue
		}

		response, err := tcpClient.Receive()
		if err != nil {
			log.Printf("Failed to receive response: %v", err)
			continue
		}

		fmt.Printf("Server response: %s\n", response)
	}
}
