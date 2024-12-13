package main

import (
	"context"
	"fmt"

	pb "github.com/TylerAldrich814/common/api"
	"go.opentelemetry.io/otel/trace"
)

type TelemetryMiddleware struct {
  next OrdersService
}

func NewTelemetryMiddleWare(
  next OrdersService,
) OrdersService {
  return &TelemetryMiddleware{ next }
}

func(t *TelemetryMiddleware) CreateOrder(
  ctx   context.Context,
  req   *pb.CreateOrderRequest,
  items []*pb.Item,
)( *pb.CreateOrderResponse, error ){
  span := trace.SpanFromContext(ctx)
  span.AddEvent(fmt.Sprintf("CreateOrder: %v", req))

  return t.next.CreateOrder(ctx, req, items)
}

func(t *TelemetryMiddleware) ValidateOrder(
  ctx context.Context,
  req *pb.CreateOrderRequest,
)( []*pb.Item, error ){
  span := trace.SpanFromContext(ctx)
  span.AddEvent(fmt.Sprintf("ValidateOrder: %v", req))

  return t.next.ValidateOrder(ctx, req)
}

func(t *TelemetryMiddleware) GetOrder(
  ctx  context.Context,
  req  *pb.GetOrderRequest,
)( *pb.CreateOrderResponse, error ){
  span := trace.SpanFromContext(ctx)
  span.AddEvent(fmt.Sprintf("GetOrder: %v", req))

  return t.next.GetOrder(ctx, req)
}

func(t *TelemetryMiddleware) UpdateOrder(
  ctx context.Context,
  req *pb.CreateOrderResponse,
)( *pb.CreateOrderResponse,error ){
  span := trace.SpanFromContext(ctx)
  span.AddEvent(fmt.Sprintf("UpdateOrder: %v", req))

  return t.next.UpdateOrder(ctx, req)
}

func(t *TelemetryMiddleware) DeleteOrder(
  ctx context.Context,
  req *pb.CreateOrderResponse,
) error {
  span := trace.SpanFromContext(ctx)
  span.AddEvent(fmt.Sprintf("DeleteOrder: %v", req))

  return t.next.DeleteOrder(ctx, req)
}
