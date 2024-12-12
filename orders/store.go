package main

import (
	"context"
	"log"

	"github.com/TylerAldrich814/common"
	pb "github.com/TylerAldrich814/common/api"
)

var orders = make([]*pb.CreateOrderResponse, 0)

type store struct {
}

func NewStore() *store {
  return &store{}
}

func(s *store) Create(
  ctx   context.Context,
  req   *pb.CreateOrderRequest,
  items []*pb.Item,
)( string, error ){
  id := "42"
  res := &pb.CreateOrderResponse{
    Id          : id,
    CustomerId  : req.CustomerId,
    Status      : "pending",
    Items       : items,
  }

  orders = append(orders, res)
  log.Printf(" ->> Stored Order #%s in local Storage\n", id)

  return id, nil
}

func(s *store) Get(
  ctx        context.Context,
  customerID string,
  id         string,
)( *pb.CreateOrderResponse,  error ){
  for _, order := range orders {
    if order.Id == id && order.CustomerId == customerID {
      return order, nil
    }
  }
  return nil, common.ErrorOrderNotFound
}

func(s *store) Update(
  ctx   context.Context,
  id    string,
  order *pb.CreateOrderResponse,
) error {
  for i, o := range orders {
    if o.Id == id {
      orders[i].Status      = order.Status
      orders[i].PaymentLink = order.PaymentLink
      return nil
    }
  }

  return common.ErrorOrderNotFound
}

func(s *store) Delete(
  ctx   context.Context,
  id    string,
  order *pb.CreateOrderResponse,
) error {
  for i, order := range orders {
    if order.Id == id {
      orders = append(orders[:i], orders[i+1:]...)
      return nil
    }
  }

  return common.ErrorOrderNotFound
}
