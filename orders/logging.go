package main

import (
	"context"

	"github.com/TylerAldrich814/common"
	pb "github.com/TylerAldrich814/common/api"
)

type LoggingMiddleware struct {
  next OrdersService
}

func NewLoggerMiddleware(
  next OrdersService,
) OrdersService {
  return &LoggingMiddleware{ next }
}

func(l *LoggingMiddleware) CreateOrder(
  ctx   context.Context,
  req   *pb.CreateOrderRequest,
  items []*pb.Item,
)( *pb.CreateOrderResponse, error ){
  defer common.LogInfo("CreateOrder")()

  return l.next.CreateOrder(ctx, req, items)
}

func(l *LoggingMiddleware) ValidateOrder(
  ctx context.Context,
  req *pb.CreateOrderRequest,
)( []*pb.Item, error ){
  defer common.LogInfo("ValidateOrder")()

  return l.next.ValidateOrder(ctx, req)
}

func(l *LoggingMiddleware) GetOrder(
  ctx  context.Context,
  req  *pb.GetOrderRequest,
)( *pb.CreateOrderResponse, error ){
  defer common.LogInfo("GetOrder")()

  return l.next.GetOrder(ctx, req)
}

func(l *LoggingMiddleware) UpdateOrder(
  ctx context.Context,
  req *pb.CreateOrderResponse,
)( *pb.CreateOrderResponse,error ){
  defer common.LogInfo("UpdateOrder")()

  return l.next.UpdateOrder(ctx, req)
}

func(l *LoggingMiddleware) DeleteOrder(
  ctx context.Context,
  req *pb.CreateOrderResponse,
) error {
  defer common.LogInfo("DeleteOrder")()

  return l.next.DeleteOrder(ctx, req)
}
