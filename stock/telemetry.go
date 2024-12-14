package main

import (
	"context"
	"fmt"

	pb "github.com/TylerAldrich814/common/api"
	"go.opentelemetry.io/otel/trace"
)

type TelemetryMiddleware struct {
  next StockService
}

func NewTelemetryMiddleware(next StockService) StockService {
  return &TelemetryMiddleware{ next }
}


func(t *TelemetryMiddleware) CheckIfItemsAreInStock(
  ctx context.Context, 
  items []*pb.ItemsWithQuantity,
)( bool,[]*pb.Item,error ) {
  span := trace.SpanFromContext(ctx)
  span.AddEvent(fmt.Sprint("CheckIfItemsAreInStock: %v", items))
  
  return t.next.CheckIfItemsAreInStock(ctx, items)
}
func(t *TelemetryMiddleware) GetItems(
  ctx context.Context, 
  ids []string,
)( []*pb.Item,error ){
  span := trace.SpanFromContext(ctx)
  span.AddEvent(fmt.Sprintf("GetItems: %v", ids))

  return t.next.GetItems(ctx, ids)
}
