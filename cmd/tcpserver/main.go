package main

import (
	"log"

	"github.com/ravidhavlesha/go-multiproto-kvs/pkg/kvstore"
	"github.com/ravidhavlesha/go-multiproto-kvs/pkg/protocol/tcp"
)

func main() {
	// Create a new KVStore
	kvStore := kvstore.NewKVStore()
	// Create a new TCP server
	tcpServer := tcp.NewTCPServer(":8080", kvStore)

	// Start listening to the new TCP server and handle connections
	if err := tcpServer.Start(); err != nil {
		log.Fatalf("Error starting TCP server: %v", err)
	}
}
