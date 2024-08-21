package main

import (
	"context"
	"log"

	pb "github.com/teclegacy/ms/oms/common/api"
	"google.golang.org/grpc"
)

type grpcHandler struct {
	pb.UnimplementedOrderServiceServer
}

func NewGrpcHandler(grpcSrv *grpc.Server) {
	grpcHandler := &grpcHandler{}
	pb.RegisterOrderServiceServer(grpcSrv, grpcHandler)
}

func (g *grpcHandler) CreateOrder(ctx context.Context, payload *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {

	log.Printf("CreateOrder request received: %v", payload)

	return &pb.CreateOrderResponse{
		Id:         "123",
		CustomerId: "123",
		Items: []*pb.ItemsWithQuantity{
			{
				ItemId:   "123",
				Quantity: 1,
			},
		},
		Status: "created",
	}, nil
}
