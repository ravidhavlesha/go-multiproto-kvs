package tcp

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"github.com/ravidhavlesha/go-multiproto-kvs/pkg/kvstore"
)

type TCPServer struct {
	address string
	kvStore *kvstore.KVStore
}

func NewTCPServer(address string, kvStore *kvstore.KVStore) *TCPServer {
	return &TCPServer{address: address, kvStore: kvStore}
}

func (server *TCPServer) Start() error {
	listner, err := net.Listen("tcp", server.address)
	if err != nil {
		return fmt.Errorf("failed to listen to tcp server: %v", err)
	}
	defer listner.Close()

	log.Printf("TCP server started on %s", server.address)

	for {
		conn, err := listner.Accept()
		if err != nil {
			log.Printf("TCP server failed to accept connection: %v", err)
			continue
		}
		go server.handleClient(conn)
	}
}

func (server *TCPServer) handleClient(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		conn.SetReadDeadline(time.Now().Add(5 * time.Minute))

		line := scanner.Text()

		response := server.handleCommand(line)
		_, err := conn.Write([]byte(response + "\n"))
		if err != nil {
			log.Printf("TCP server failed to write data to client: %v", err)
			return
		}
	}
}

func (server *TCPServer) handleCommand(line string) string {
	parts := strings.Fields(line)

	if len(parts) < 1 {
		return "Error: Invalid command"
	}

	command := strings.ToUpper(parts[0])

	switch command {
	case "GET":
		if len(parts) != 2 {
			return "Usage: GET <key>"
		}
		key := parts[1]
		value, exists := server.kvStore.Get(key)
		if !exists {
			return "Error: Key not found"
		}
		return value

	case "SET":
		if len(parts) != 3 {
			return "Usage: SET <key> <value>"
		}
		key := parts[1]
		value := parts[2]
		server.kvStore.Set(key, value)
		return "OK"

	case "DELETE":
		if len(parts) != 2 {
			return "Usage: DELETE <key>"
		}
		key := parts[1]
		server.kvStore.Delete(key)
		return "OK"

	default:
		return "Error: Unknown command"
	}

}
