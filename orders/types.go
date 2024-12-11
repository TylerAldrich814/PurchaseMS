package main

import (
  "context"
	pb "github.com/TylerAldrich814/common/api"
)

type OrdersService interface {
  CreateOrder(context.Context,  *pb.CreateOrderRequest, []*pb.Item)( *pb.CreateOrderResponse, error )
  ValidateOrder(context.Context, *pb.CreateOrderRequest)( []*pb.Item, error )
  GetOrder(context.Context,  *pb.GetOrderRequest)( *pb.CreateOrderResponse, error )
}

type OrdersStore interface {
  Create(context.Context, *pb.CreateOrderRequest, []*pb.Item)(string, error)
  Get(ctx context.Context, customerID, id string)( *pb.CreateOrderResponse, error )
}
