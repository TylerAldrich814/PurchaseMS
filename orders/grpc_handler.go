package main

import (
	"context"
	"log"

	pb "github.com/TylerAldrich814/common/api"
	"google.golang.org/grpc"
)

type grpcHandler struct {
  pb.UnimplementedOrderServiceServer
}

func NewGRPCHandler(grpcServer *grpc.Server) {
 handler := &grpcHandler{}
 pb.RegisterOrderServiceServer(grpcServer, handler)
}

func(grpc *grpcHandler) CreateOrder(
  ctx context.Context, 
  r   *pb.CreateOrderRequest,
) (*pb.CreateOrderResponse, error) {
  log.Println("New Order Received")
  order := &pb.CreateOrderResponse{
    Id: "42",
  }

  return order, nil
}
