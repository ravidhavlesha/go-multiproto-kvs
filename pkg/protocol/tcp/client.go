package tcp

import (
	"bufio"
	"fmt"
	"net"
)

type TCPClient struct {
	address string
	conn    net.Conn
	scanner *bufio.Scanner
}

func NewTCPClient(address string) *TCPClient {
	return &TCPClient{address: address}
}

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

func (client *TCPClient) Receive() (string, error) {
	if client.scanner == nil {
		return "", fmt.Errorf("scanner is not initialized")
	}

	if client.scanner.Scan() {
		return client.scanner.Text(), nil
	}

	if err := client.scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading response from server: %w", err)
	}

	return "", fmt.Errorf("server closed the connection")
}

func (client *TCPClient) Close() error {
	if client.conn != nil {
		return client.conn.Close()
	}
	return nil
}
