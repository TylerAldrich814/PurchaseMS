package gateway

import (
  "context"

	pb "github.com/TylerAldrich814/common/api"
)

type OrdersGateway interface {
  CreateOrder(context.Context, *pb.CreateOrderRequest)( *pb.CreateOrderResponse, error )
}
