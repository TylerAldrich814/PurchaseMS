package main

import (
  "context"
	pb "github.com/TylerAldrich814/common/api"
)

type OrdersService interface {
  CreateOrder(context.Context,  *pb.CreateOrderRequest)( *pb.CreateOrderResponse, error )
  ValidateOrder(context.Context, *pb.CreateOrderRequest)( []*pb.Item, error )
}

type OrdersStore interface {
  Create(context.Context) error
}
