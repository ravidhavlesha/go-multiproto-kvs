package grpc

import (
	"context"
	"fmt"
	"time"

	pb "github.com/ravidhavlesha/go-multiproto-kvs/internal/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type KVStoreClient struct {
	conn   *grpc.ClientConn
	client pb.KVStoreClient
}

func NewKVStoreClient(address string) (*KVStoreClient, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to gRPC server: %w", err)
	}

	client := pb.NewKVStoreClient(conn)
	return &KVStoreClient{conn: conn, client: client}, nil
}

func (client *KVStoreClient) Get(key string) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	resp, err := client.client.Get(ctx, &pb.GetRequest{Key: key})
	if err != nil {
		return "", false, fmt.Errorf("failed to GET key: %w", err)
	}

	return resp.GetValue(), resp.GetFound(), nil
}

func (client *KVStoreClient) Set(key, value string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	_, err := client.client.Set(ctx, &pb.SetRequest{Key: key, Value: value})
	if err != nil {
		return fmt.Errorf("failed to SET key: %w", err)
	}

	return nil
}

func (client *KVStoreClient) Delete(key string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	_, err := client.client.Delete(ctx, &pb.DeleteRequest{Key: key})
	if err != nil {
		return fmt.Errorf("failed to DELETE key: %w", err)
	}

	return nil
}

func (client *KVStoreClient) Close() error {
	return client.conn.Close()
}
