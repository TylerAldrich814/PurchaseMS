package main

import (
	"context"

	pb "github.com/TylerAldrich814/common/api"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
)

type StockGrpcHandler struct {
  pb.UnimplementedStockServiceServer

  service StockService
  channel *amqp.Channel
}

func NewGRPCHandler(
  server       *grpc.Server,
  channel      *amqp.Channel,
  stockService StockService,
){
  handler := &StockGrpcHandler{
    service : stockService,
    channel : channel,
  }

  pb.RegisterStockServiceServer(server, handler)
}

func(s *StockGrpcHandler) CheckIfItemIsInStock(
  ctx context.Context,
  req *pb.CheckIfItemIsInStockRequest,
)( *pb.CheckIfItemIsInStockResponse, error ){
  inStock, items, err := s.service.CheckIfItemsAreInStock(ctx, req.Items)
  if err != nil {
    return nil, err
  }
  return &pb.CheckIfItemIsInStockResponse{
    InStock : inStock,
    Items   : items,
  }, nil
}

func(s *StockGrpcHandler) GetItems(
  ctx context.Context,
  req *pb.GetItemsRequest,
)( *pb.GetItemsResponse, error ){
  items, err := s.service.GetItems(ctx, req.ItemIds)
  if err != nil {
    return nil, err
  }
  return &pb.GetItemsResponse{
    Items : items,
  }, nil
}
