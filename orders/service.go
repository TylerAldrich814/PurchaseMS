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

func(s *service) CreateOrder(ctx context.Context) error {
  return nil
}

func(s *service) ValidateOrder(
  ctx   context.Context,
  order *pb.CreateOrderRequest,
) error {
  if len(order.Items) == 0 {
    return common.ErrorNoItems
  }

  mergedItems := mergeItemsQuantities(order.Items)
  log.Print(mergedItems)

  // ->> Validate with the Stock Service

  return nil
}

func mergeItemsQuantities(items []*pb.ItemsWithQuantity) []*pb.ItemsWithQuantity {
  processed := make(map[string]*pb.ItemsWithQuantity)
  merged    := make([]*pb.ItemsWithQuantity, 0)

  for _, item := range items {
    if i, ok := processed[item.Id]; !ok {
      processed[i.Id] = i
      merged = append(merged, processed[i.Id])
    } else {
      processed[i.Id].Quantity += item.Quantity
    }
  }

  return merged
}
