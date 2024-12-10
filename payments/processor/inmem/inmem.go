package inmem

import (
	pb "github.com/TylerAldrich814/common/api"
	_ "github.com/TylerAldrich814/payments/processor"
)

type InMemory struct {}

func NewInMemory() *InMemory {
  return &InMemory{}
}

func(i *InMemory) CreatePaymentLink(*pb.CreateOrderResponse)( string,error ){
  return "DummyLink", nil
}
