package gateway

import (
	"context"
	"log"

	pb "github.com/TylerAldrich814/common/api"
	"github.com/TylerAldrich814/common/discovery"
)


type gateway struct {
  registry discovery.Registry
}

func NewGateway(registry discovery.Registry) *gateway {
  return &gateway{ registry }
}

func(g *gateway) UpdateOrderAfterPayment(
  ctx         context.Context, 
  orderID     string, 
  paymentLink string,
) error {
  conn, err := discovery.ServiceConnection(
    context.Background(),
    "orders",
    g.registry,
  )
  if err != nil {
    return err
  }
  defer conn.Close()

  ordersclient := pb.NewOrderServiceClient(conn)

  _, err = ordersclient.UpdateOrder(ctx, &pb.CreateOrderResponse{
    Id          :  orderID,
    Status      : "waiting_payment",
    PaymentLink : paymentLink,
  })
  log.Printf("Updated Payment: %s", orderID)
  return err
}
