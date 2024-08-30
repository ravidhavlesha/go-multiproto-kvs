package main

import (
	"fmt"
	"log"

	"github.com/ravidhavlesha/go-multiproto-kvs/pkg/protocol/grpc"
)

func main() {
	address := "localhost:50051"
	client, err := grpc.NewKVStoreClient(address)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	if err := client.Set("foo", "bar"); err != nil {
		log.Fatalf("failed to SET: %v", err)
	}

	value, found, err := client.Get("foo")
	if err != nil {
		log.Fatalf("failed to GET: %v", err)
	}
	if found {
		fmt.Printf("Found key 'foo' with value: %s\n", value)
	} else {
		fmt.Println("Key 'foo' not found")
	}

	if err := client.Delete("foo"); err != nil {
		log.Fatalf("failed to DELETE: %v", err)
	}
}
