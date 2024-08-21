package main

import (
	"context"
	"log"
	"net"

	"github.com/teclegacy/ms/oms/common"
	"google.golang.org/grpc"
)

var (
	grpcAddr = common.EnvString("GRPC_ADDR", "localhost:2000")
)

func main() {

	// tcp server
	l, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer l.Close()

	// grpc server & handler
	grpcSrv := grpc.NewServer()
	NewGrpcHandler(grpcSrv)

	store := NewStore()
	srv := NewService(store)

	srv.CreateOrder(context.Background())

	// start grpc server
	log.Printf("Server started on port %s", grpcAddr)
	if err := grpcSrv.Serve(l); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
