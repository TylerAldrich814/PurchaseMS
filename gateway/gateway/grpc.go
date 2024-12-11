package gateway

import (
	"context"

	pb "github.com/TylerAldrich814/common/api"
	"github.com/TylerAldrich814/common/discovery"
)

type gateway struct {
  registry     discovery.Registry
  ordersClient pb.OrderServiceClient
}

func NewGRPCGateway(
  ctx      context.Context,
  registry discovery.Registry,
)( *gateway, error ){
  // ->> Connect
  conn, err := discovery.ServiceConnection(ctx, "orders", registry)
  if err != nil {
    return nil, err
  }

  ordersClient := pb.NewOrderServiceClient(conn)

  return &gateway{ 
    registry,
    ordersClient,
  }, nil
}

func(g *gateway) CreateOrder(
  ctx context.Context, 
  req *pb.CreateOrderRequest,
)( *pb.CreateOrderResponse, error ){
  return g.ordersClient.CreateOrder(ctx, &pb.CreateOrderRequest{
    CustomerId : req.CustomerId,
    Items      : req.Items,
  })
}

func(g *gateway) GetOrder(
  ctx        context.Context, 
  orderID    string, 
  custonerID string,
)( *pb.CreateOrderResponse, error ){
  return g.ordersClient.GetOrder(ctx, &pb.GetOrderRequest{
    OrderId    : orderID,
    CustomerId : custonerID,
  })
}
