package main

import (
	"context"

	"github.com/TylerAldrich814/common"
	pb "github.com/TylerAldrich814/common/api"
)

type LoggingMiddleware struct {
  next StockService
}

func NewLoggingMiddleware(next StockService) StockService {
  return &LoggingMiddleware{ next }
}

func(l *LoggingMiddleware) CheckIfItemsAreInStock(
  ctx context.Context,
  items []*pb.ItemsWithQuantity,
)( bool,[]*pb.Item,error ) {
  defer common.LogInfo("CheckIfItemsAreInStock")()

  return l.next.CheckIfItemsAreInStock(ctx, items)
}
func(l *LoggingMiddleware) GetItems(
  ctx context.Context, 
  ids []string,
)( []*pb.Item,error ) {
  defer common.LogInfo("GetItems")()

  return l.next.GetItems(ctx, ids)
}
