package tcp

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"github.com/ravidhavlesha/go-multiproto-kvs/pkg/kvstore"
)

// TCPServer represents a TCP server.
type TCPServer struct {
	address string
	kvStore *kvstore.KVStore
}

// NewTCPServer initializes a new TCP server
func NewTCPServer(address string, kvStore *kvstore.KVStore) *TCPServer {
	return &TCPServer{address: address, kvStore: kvStore}
}

// Start starts a TCP server that listens to incoming connections.
func (server *TCPServer) Start() error {
	listner, err := net.Listen("tcp", server.address)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", server.address, err)
	}
	defer listner.Close()

	log.Printf("TCP server started on %s", server.address)

	for {
		conn, err := listner.Accept()
		if err != nil {
			log.Printf("failed to accept connection: %v", err)
			continue
		}
		log.Printf("Client connected from %s", conn.RemoteAddr().String())

		// Handle the connection in a new goroutine for concurrency.
		go server.handleClient(conn)
	}
}

// handleClient manages a single client connection.
func (server *TCPServer) handleClient(conn net.Conn) {
	defer func() {
		log.Printf("Client disconnected from %s", conn.RemoteAddr().String())
		conn.Close()
	}()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		// Set a read deadline to avoid blocking forever
		conn.SetReadDeadline(time.Now().Add(5 * time.Minute))

		line := scanner.Text()

		// Handle the message and respond to the client
		response, err := server.handleCommand(line)
		if err != nil {
			log.Printf("command handling error: %v", err)
			_, _ = conn.Write([]byte("Error: " + err.Error() + "\n"))
			continue
		}

		_, err = conn.Write([]byte(response + "\n"))
		if err != nil {
			log.Printf("failed to write response to client: %v", err)
			continue
		}
	}

	// Log the error if the scanner encounters one.
	if err := scanner.Err(); err != nil {
		log.Printf("connection error: %v", err)
	}
}

// handleCommand processes a client command and interacts with the KVStore.
func (server *TCPServer) handleCommand(line string) (string, error) {
	parts := strings.SplitN(line, " ", 3)

	if len(parts) < 1 {
		return "", errors.New("invalid command")
	}

	command := strings.ToUpper(parts[0])

	switch command {
	case "GET":
		if len(parts) != 2 {
			return "", errors.New("usage: GET <key>")
		}
		value, exists := server.kvStore.Get(parts[1])
		if !exists {
			return "", errors.New("key not found")
		}
		return value, nil

	case "SET":
		if len(parts) != 3 {
			return "", errors.New("usage: SET <key> <value>")
		}
		if err := server.kvStore.Set(parts[1], parts[2]); err != nil {
			return "", fmt.Errorf("failed to set value: %w", err)
		}
		return "OK", nil

	case "DELETE":
		if len(parts) != 2 {
			return "", errors.New("usage: DELETE <key>")
		}

		if err := server.kvStore.Delete(parts[1]); err != nil {
			return "", fmt.Errorf("failed to delete key: %w", err)
		}
		return "OK", nil

	default:
		return "", errors.New("unknown command")
	}

}
