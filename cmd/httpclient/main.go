package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ravidhavlesha/go-multiproto-kvs/pkg/protocol/http"
)

func main() {
	// Create a new HTTP client
	httpClient := http.NewHTTPClient("http://localhost:8080")

	// Create a reader for standard input (for user commands)
	reader := bufio.NewReader(os.Stdin)

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

		parts := strings.SplitN(userInput, " ", 3)
		if len(parts) == 0 {
			fmt.Println("Invalid command")
			continue
		}

		command := strings.ToUpper(parts[0])
		switch command {
		case "GET":
			if len(parts) != 2 {
				fmt.Println("Usage: GET <key>")
				continue
			}
			value, err := httpClient.Get(parts[1])
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			} else {
				fmt.Printf("Value: %s\n", value)
			}

		case "SET":
			if len(parts) != 3 {
				fmt.Println("Usage: SET <key> <value>")
				continue
			}
			if err := httpClient.Set(parts[1], parts[2]); err != nil {
				fmt.Printf("Error: %v\n", err)
			} else {
				fmt.Println("OK")
			}

		case "DELETE":
			if len(parts) != 2 {
				fmt.Println("Usage: DELETE <key>")
				continue
			}
			if err := httpClient.Delete(parts[1]); err != nil {
				fmt.Printf("Error: %v\n", err)
			} else {
				fmt.Println("OK")
			}

		default:
			fmt.Println("Unknown command")
		}
	}
}
