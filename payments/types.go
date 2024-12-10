package main

import (
  "context"

	pb "github.com/TylerAldrich814/common/api"
)

type PaymentService interface {
  CreatePayments(context.Context, *pb.CreateOrderResponse)( string,error )
}
