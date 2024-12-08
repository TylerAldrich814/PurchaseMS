package main

import (
	"context"
	"errors"
	"log"

	pb "github.com/TylerAldrich814/common/api"
	"google.golang.org/grpc"
)

type grpcHandler struct {
  pb.UnimplementedOrderServiceServer

  service OrdersService
}

func NewGRPCHandler(
  grpcServer *grpc.Server,
  service     OrdersService,
) {
 handler := &grpcHandler{
   service : service,
 }
 pb.RegisterOrderServiceServer(grpcServer, handler)
}

func(grpc *grpcHandler) CreateOrder(
  ctx context.Context, 
  r   *pb.CreateOrderRequest,
) (*pb.CreateOrderResponse, error) {
  log.Printf("New Order Received: %s", r)


  order := &pb.CreateOrderResponse{
    Id: "42",
  }

  return order, nil
}

func(grpc *grpcHandler) ValidateOrder(
  ctx   context.Context, 
  order *pb.CreateOrderRequest,
) error {
  if len(order.Items) == 0 {
    return errors.New("")
  }

  return nil
}
