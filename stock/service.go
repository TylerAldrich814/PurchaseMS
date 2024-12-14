package main

import (
  "context"
  pb "github.com/TylerAldrich814/common/api"
)

type Service struct {
  store StockStore
}

func NewService(store StockStore) *Service {
  return &Service{ store }
}

func(s *Service) CheckIfItemsAreInStock(
  ctx context.Context, 
  items []*pb.ItemsWithQuantity,
)( bool,[]*pb.Item,error ) {
  itemIDs := make([]string, 0)
  for _, item := range items {
    itemIDs = append(itemIDs, item.Id)
  }
  itemsInStock, err := s.store.GetItems(ctx, itemIDs)
  if err != nil {
    return false, nil, err
  }

  for _, stockItem := range itemsInStock {
    for _, reqItems := range items {
      if stockItem.Id == reqItems.Id && stockItem.Quantity < reqItems.Quantity {
        return false, itemsInStock, nil
      }
    }
  }
  stockItems := make([]*pb.Item, 0)
  for _, stockItem := range itemsInStock {
    for _, reqItem := range items {
      if stockItem.Id == reqItem.Id {
        stockItems = append(stockItems, &pb.Item {
          Id       : stockItem.Id,
          Name     : stockItem.Name,
          PriceId  : stockItem.PriceId,
          Quantity : stockItem.Quantity,
        })
      }
    }
  }

  return true, stockItems, nil
}

func(s *Service) GetItems(
  ctx context.Context, 
  ids []string,
)( []*pb.Item,error ){
  return s.store.GetItems(ctx, ids)
}
