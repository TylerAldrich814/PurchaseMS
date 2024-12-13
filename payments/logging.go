package main

import (
	"context"
	"github.com/TylerAldrich814/common"
	pb "github.com/TylerAldrich814/common/api"
)

type LoggingMiddleware struct {
  next PaymentService
}

func NewLoggingMiddleware(
  next PaymentService,
) PaymentService {
  return &LoggingMiddleware{ next }
}

func(l *LoggingMiddleware) CreatePayments(
  ctx context.Context,
  req *pb.CreateOrderResponse,
)( string,error ){
  defer common.LogInfo("CreatePayments")()

  return l.next.CreatePayments(ctx, req)
}
