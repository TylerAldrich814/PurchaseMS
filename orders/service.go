package main

import ( "context"
	"log"

	pb "github.com/TylerAldrich814/common/api"
	"github.com/TylerAldrich814/common/errors"
	"github.com/TylerAldrich814/orders/gateway"
)

type service struct {
  store   OrdersStore
  gateway gateway.StockGateway
}

func NewService(
  store   OrdersStore,
  gateway gateway.StockGateway,
) *service {
  return &service{
    store,
    gateway,
  }
}

func(s *service) CreateOrder(
  ctx   context.Context,
  req   *pb.CreateOrderRequest,
  items []*pb.Item,
)( *pb.CreateOrderResponse, error ){
  id, err := s.store.Create(ctx, req, items)
  if err != nil {
    return nil, err
  }

  order := &pb.CreateOrderResponse{
    Id: id,
    CustomerId: req.CustomerId,
    Status: "pending",
    Items: items,
  }
  return order, nil
}

func(s *service) ValidateOrder(
  ctx context.Context,
  req *pb.CreateOrderRequest,
)( []*pb.Item, error ){
  if len(req.Items) == 0 {
    return nil, errors.NoItems
  }

  // ->> Validate with the Stock Service
  inStock, items, err := s.gateway.CheckIfItemIsInStock(
    ctx,
    req.CustomerId,
    mergeItemsQuantities(req.Items),
  )
  if err != nil {
    return nil, err
  }
  if !inStock {
    return nil, errors.Stock_ItemOutOfStock
  }

  return items, nil
}

func(s *service) GetOrder(
  ctx  context.Context,
  req  *pb.GetOrderRequest,
)( *pb.CreateOrderResponse, error ){
  return s.store.Get(
    ctx,
    req.OrderId,
    req.CustomerId,
  )
}

func(s *service) UpdateOrder(
  ctx context.Context,
  req *pb.CreateOrderResponse,
)( *pb.CreateOrderResponse,error ){
  if err := s.store.Update(
    ctx,
    req.Id,
    req,
  ); err != nil {
    return nil, err
  }

  return req, nil
}

func(s *service) DeleteOrder(
  ctx context.Context,
  req *pb.CreateOrderResponse,
) error {

  return nil
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
