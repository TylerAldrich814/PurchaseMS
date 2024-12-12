package main

import (
  "log"
	"context"

	pb "github.com/TylerAldrich814/common/api"
	devtools "github.com/TylerAldrich814/common/dev_tools"
	"github.com/TylerAldrich814/payments/gateway"
	"github.com/TylerAldrich814/payments/processor"
)

type service struct {
  processor processor.PaymentProcessor
  gateway   gateway.PaymentsGateway
}

func NewService(
  processor processor.PaymentProcessor,
  gateway   gateway.PaymentsGateway,
) *service {
  return &service{ 
    processor,
    gateway,
  }
}

func(s *service) CreatePayments(
  ctx context.Context,
  req *pb.CreateOrderResponse,
)( string,error ){
  link, err := s.processor.CreatePaymentLink(req)
  if err != nil {
    return "", err
  }
  if err := devtools.ClipboardCopy(link); err != nil {
    log.Printf(err.Error())
  }

  // ->> Update Order with Forward Link
  if err := s.gateway.UpdateOrderAfterPayment(
    ctx,
    req.Id,
    link,
  ); err != nil {
    return "", err
  }

  return link, nil
}
