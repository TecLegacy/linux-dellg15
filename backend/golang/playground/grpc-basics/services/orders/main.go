package main

import "log"

const (
	addr = "localhost:8080"
)

func main() {
	// Create a new gRPC server
	g := NewGrpcServer(addr)
	if err := g.Run(); err != nil {
		log.Fatalf("failed to run gRPC server: %v", err)
	}

	// Create a new HTTP server

}
