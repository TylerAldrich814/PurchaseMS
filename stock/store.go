package main

import (
	"context"

	pb "github.com/TylerAldrich814/common/api"
	"github.com/TylerAldrich814/common/errors"
)

type Store struct {
  stock map[string]*pb.Item
}

func NewStore() *Store {
  return &Store{
    stock: map[string]*pb.Item {
      "2": {
        Id       :"2",
        Name     :"Coke",
        PriceId  : "price_1QVzRuGzCrbvBMvKDd88FGDv",
        Quantity : 10,
      },
      "1": {
        Id       : "1",
        Name     : "Cheese",
        PriceId  : "price_1QUDOtGzCrbvBMvK9wsNawMv",
        Quantity : 20,
      },
    },
  }
}

func(s *Store) GetItem(
  ctx context.Context,
  id  string,
)( *pb.Item, error) {
  for _, item := range s.stock {
    if item.Id == id {
      return item, nil
    }
  }
  return nil, errors.Stock_ItemNotFound
}

func(s *Store) GetItems(
  ctx context.Context,
  ids []string,
)( []*pb.Item, error ){
  var res []*pb.Item
  for _, id := range ids {
    if i, ok := s.stock[id]; ok {
      res = append(res, i)
    }
  }
  return res, nil
}
