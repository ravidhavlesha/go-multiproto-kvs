package grpc

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/ravidhavlesha/go-multiproto-kvs/internal/proto"
	"github.com/ravidhavlesha/go-multiproto-kvs/pkg/kvstore"
	"google.golang.org/grpc"
)

type KVStoreServer struct {
	kvstore *kvstore.KVStore
	pb.UnimplementedKVStoreServer
}

func NewKVStoreServer(store *kvstore.KVStore) *KVStoreServer {
	return &KVStoreServer{kvstore: store}
}

func (kvs *KVStoreServer) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	value, found := kvs.kvstore.Get(req.GetKey())
	return &pb.GetResponse{Value: value, Found: found}, nil
}

func (kvs *KVStoreServer) Set(ctx context.Context, req *pb.SetRequest) (*pb.SetResponse, error) {
	kvs.kvstore.Set(req.GetKey(), req.GetValue())
	return &pb.SetResponse{Status: "OK"}, nil
}

func (kvs *KVStoreServer) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	kvs.kvstore.Delete(req.GetKey())
	return &pb.DeleteResponse{Status: "OK"}, nil
}

func StartGRPCServer(address string, store *kvstore.KVStore) error {
	listner, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", address, err)
	}
	defer listner.Close()

	server := grpc.NewServer()
	pb.RegisterKVStoreServer(server, NewKVStoreServer(store))

	log.Printf("GRPC server started on %s", address)
	return server.Serve(listner)
}
