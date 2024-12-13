package main

import (
	"context"
	"fmt"

	pb "github.com/TylerAldrich814/common/api"
	"go.opentelemetry.io/otel/trace"
)

type TelemetryMiddleware struct {
  next PaymentService
}

func NewTelemetryMiddleware(
  next PaymentService,
) *TelemetryMiddleware {
  return &TelemetryMiddleware{ next }
}

func(t *TelemetryMiddleware) CreatePayments(
  ctx context.Context,
  req *pb.CreateOrderResponse,
)( string,error ){
  span := trace.SpanFromContext(ctx)
  span.AddEvent(fmt.Sprintf("CreatePayments: %v", req))

  return t.next.CreatePayments(ctx, req)
}
