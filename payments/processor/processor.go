package processor

import (
	pb "github.com/TylerAldrich814/common/api"
)

type PaymentProcessor interface {
  CreatePaymentLink(*pb.CreateOrderResponse)( string,error )
}
