package main

import (
	"log"

	"github.com/ravidhavlesha/go-multiproto-kvs/pkg/kvstore"
	"github.com/ravidhavlesha/go-multiproto-kvs/pkg/protocol/http"
)

func main() {
	// Create a new KVStore
	kvStore := kvstore.NewKVStore()
	// Create a new HTTP server
	httpServer := http.NewHTTPServer(":8080", kvStore)

	// Start listening to the new HTTP server and handle connections
	if err := httpServer.Start(); err != nil {
		log.Fatalf("Error starting HTTP server: %v", err)
	}
}
