package main

import (
	"context"

	pb "github.com/TylerAldrich814/common/api"
	"github.com/TylerAldrich814/payments/processor"
)

type service struct {
  processor processor.PaymentProcessor
}

func NewService(
  processor processor.PaymentProcessor,
) *service {
  return &service{ processor }
}

func(s *service) CreatePayments(
  ctx context.Context,
  req *pb.CreateOrderResponse,
)( string,error ){
  link, err := s.processor.CreatePaymentLink(req)
  if err != nil {
    return "", err
  }

  return link, nil
}
