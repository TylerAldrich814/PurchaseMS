package gateway

import (
  "context"
  pb "github.com/TylerAldrich814/common/api"
)

type StockGateway interface {
  CheckIfItemIsInStock(
    ctx        context.Context, 
    customerId string, 
    items      []*pb.ItemsWithQuantity,
  )( bool, []*pb.Item, error )
}
