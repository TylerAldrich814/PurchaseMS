package gateway

import (
	"context"
	"log"

	pb "github.com/TylerAldrich814/common/api"
	"github.com/TylerAldrich814/common/discovery"
)

type Gateway struct {
  registry discovery.Registry
}

func NewGateway(registry discovery.Registry) *Gateway {
  return &Gateway{ registry }
}

func(g *Gateway) CheckIfItemIsInStock(
  ctx        context.Context, 
  customerId string, 
  items      []*pb.ItemsWithQuantity,
)( bool, []*pb.Item, error ){
  conn, err := discovery.ServiceConnection(
    context.Background(),
    "stock",
    g.registry,
  )
  if err != nil {
    log.Fatalf("Failed to dial Server: %v", err)
  }
  defer conn.Close()

  client := pb.NewStockServiceClient(conn)

  res, err := client.CheckIfItemIsInStock(
    ctx, 
    &pb.CheckIfItemIsInStockRequest{
      Items: items,
    },
  )

  return res.InStock, res.Items, err
}