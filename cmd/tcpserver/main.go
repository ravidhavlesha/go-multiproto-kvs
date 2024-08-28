package main

import (
	"log"

	"github.com/ravidhavlesha/go-multiproto-kvs/pkg/kvstore"
	"github.com/ravidhavlesha/go-multiproto-kvs/pkg/protocol/tcp"
)

func main() {
	kvStore := kvstore.NewKVStore()
	tcpServer := tcp.NewTCPServer(":8080", kvStore)

	if err := tcpServer.Start(); err != nil {
		log.Fatalf("Error starting TCP server: %v", err)
	}
}
