package gateway

import (
	"context"

	pb "github.com/TylerAldrich814/common/api"
	"github.com/TylerAldrich814/common/discovery"
)

type gateway struct {
  registry discovery.Registry
}

func NewGRPCGateway(registry discovery.Registry) *gateway {
  return &gateway{ registry }
}

func(g *gateway) CreateOrder(
  ctx context.Context, 
  req *pb.CreateOrderRequest,
)( *pb.CreateOrderResponse, error ){
  // ->> Connect
  conn, err := discovery.ServiceConnection(ctx, "orders", g.registry)
  if err != nil {
    return nil, err
  }

  client := pb.NewOrderServiceClient(conn)

  return client.CreateOrder(ctx, &pb.CreateOrderRequest{
    CustomerId : req.CustomerId,
    Items      : req.Items,
  })
}
