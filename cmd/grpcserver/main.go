package main

import (
	"log"

	"github.com/ravidhavlesha/go-multiproto-kvs/pkg/kvstore"
	"github.com/ravidhavlesha/go-multiproto-kvs/pkg/protocol/grpc"
)

func main() {
	// Create a new KVStore
	kvStore := kvstore.NewKVStore()

	// Start listening to the new GRPC server and handle connections
	if err := grpc.StartGRPCServer(":50051", kvStore); err != nil {
		log.Fatalf("Error starting GRPC server: %v", err)
	}
}
