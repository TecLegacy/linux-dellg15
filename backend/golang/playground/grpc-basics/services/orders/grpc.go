package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

type grpcServer struct {
	addr string
}

func NewGrpcServer(addr string) *grpcServer {
	return &grpcServer{
		addr,
	}
}

func (s *grpcServer) Run() error {
	l, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	gSrv := grpc.NewServer()

	// Register the service with the server

	log.Printf("gRPC server listening on %s", s.addr)
	return gSrv.Serve(l)

}
