package main

import (
	"context"
	"log"

	"github.com/TylerAldrich814/common"
	pb "github.com/TylerAldrich814/common/api"
)

type service struct {
  store OrdersStore
}

func NewService(store OrdersStore) *service {
  return &service{store}
}

func(s *service) CreateOrder(
  ctx context.Context,
  req *pb.CreateOrderRequest,
)( *pb.CreateOrderResponse, error ){
  items, err := s.ValidateOrder(ctx, req)
  if err != nil {
    return nil, err
  }

  order := &pb.CreateOrderResponse {
    Id         : "42",
    CustomerId : req.CustomerId,
    Status     : "pending",
    Items      : items,
  }

  return order, err
}

func(s *service) ValidateOrder(
  ctx context.Context,
  req *pb.CreateOrderRequest,
)( []*pb.Item, error ){
  if len(req.Items) == 0 {
    return nil, common.ErrorNoItems
  }

  mergedItems := mergeItemsQuantities(req.Items)
  log.Print(mergedItems)

  // ->> Validate with the Stock Service

  // TODO: TEMP
  var itemsWithPriceTemp []*pb.Item
  for _, item := range mergedItems {
    itemsWithPriceTemp = append(itemsWithPriceTemp, &pb.Item{
      PriceId  : "price_1QUDOtGzCrbvBMvK9wsNawMv",
      Id       : item.Id,
      Quantity : item.Quantity,
    })
  }

  return itemsWithPriceTemp, nil
}

func mergeItemsQuantities(
  items []*pb.ItemsWithQuantity,
) []*pb.ItemsWithQuantity {
  processed := make(map[string]*pb.ItemsWithQuantity)
  merged    := make([]*pb.ItemsWithQuantity, 0)

  for _, item := range items {
    log.Printf("Merging Item %s\n", item.Id)
    if i, ok := processed[item.Id]; !ok {
      processed[item.Id] = item
      merged = append(merged, processed[item.Id])
    } else {
      processed[i.Id].Quantity += item.Quantity
    }
  }

  return merged
}
