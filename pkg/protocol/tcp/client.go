package tcp

import (
	"bufio"
	"fmt"
	"net"
)

// TCPClient represents a TCP client.
type TCPClient struct {
	address string
	conn    net.Conn
	scanner *bufio.Scanner
}

// NewTCPClient initializes a new TCP client
func NewTCPClient(address string) *TCPClient {
	return &TCPClient{address: address}
}

// Connect establishes to the TCP server
func (client *TCPClient) Connect() error {
	conn, err := net.Dial("tcp", client.address)
	if err != nil {
		return fmt.Errorf("failed to connect to TCP server: %v", err)
	}

	client.conn = conn
	client.scanner = bufio.NewScanner(conn)
	fmt.Println("Connected to the server.")
	return nil
}

// Send sends a command to the TCP server
func (client *TCPClient) Send(cmd string) error {
	if client.conn == nil {
		return fmt.Errorf("connection is not established")
	}

	_, err := fmt.Fprintf(client.conn, cmd+"\n")
	if err != nil {
		return fmt.Errorf("failed to send command: %w", err)
	}
	return nil
}

// Receive reads and return the response from the TCP server
func (client *TCPClient) Receive() (string, error) {
	if client.scanner == nil {
		return "", fmt.Errorf("scanner is not initialized")
	}

	if client.scanner.Scan() {
		return client.scanner.Text(), nil
	}

	// Log the error if the scanner encounters one.
	if err := client.scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading response from server: %w", err)
	}

	return "", fmt.Errorf("server closed the connection")
}

// Close closes the connection to the TCP server
func (client *TCPClient) Close() error {
	if client.conn != nil {
		return client.conn.Close()
	}
	return nil
}
