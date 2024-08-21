package main

import (
	"log"
	"net/http"

	_ "github.com/joho/godotenv/autoload"
	"github.com/teclegacy/ms/oms/common"
	pb "github.com/teclegacy/ms/oms/common/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	httpAddr = common.EnvString("HTTP_ADDR", ":8080")

	// TODO: fix : Make sure its listening to more than one instance of order service
	addrToOrderService = "localhost:3000"
)

func main() {

	// Establish grpc channel connection
	conn, err := grpc.NewClient(addrToOrderService, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to dial server %v", err)
	}
	defer conn.Close()

	// Create a new instance of order service client
	c := pb.NewOrderServiceClient(conn)

	mux := http.NewServeMux()

	handler := NewHandler(c)
	handler.RegisterRoutes(mux)

	log.Printf("Server started on port %s", httpAddr)
	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatal("Failed to start server")
	}

}
