package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	pb "github.com/TylerAldrich814/common/api"
	"github.com/TylerAldrich814/common/broker"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.opentelemetry.io/otel"
	"google.golang.org/grpc"
)

type grpcHandler struct {
  pb.UnimplementedOrderServiceServer
  service OrdersService
  channel *amqp.Channel
}

func NewGRPCHandler(
  grpcServer *grpc.Server,
  service     OrdersService,
  channel    *amqp.Channel,
) {
 handler := &grpcHandler{
   service : service,
   channel : channel,
 }
 pb.RegisterOrderServiceServer(grpcServer, handler)
}

func(grpc *grpcHandler) CreateOrder(
  ctx context.Context, 
  req   *pb.CreateOrderRequest,
) (*pb.CreateOrderResponse, error) {
  // ->> Communicate to the Broker that a new Order has been recorded
  q, err := grpc.channel.QueueDeclare(
    broker.OrderCreatedEvent,
    true,
    false,
    false,
    false,
    nil,
  )
  if err != nil {
    log.Fatal(err)
  }

  tr := otel.Tracer("amqp")
  amqpContext, msgSpan := tr.Start(
    ctx,
    fmt.Sprintf(
      "AMQP - publish - %s",
      q.Name,
    ),
  )
  defer msgSpan.End()

  // ->> Validate incoming Order(s)
  items, err := grpc.service.ValidateOrder(amqpContext, req)
  if err != nil {
    return nil, err
  }
  
  // ->> Store Incoming Order(s) in Order Service DB
  order, err := grpc.service.CreateOrder(amqpContext, req, items)
  if err != nil {
    return nil, err
  }

  marshalledOrder, err := json.Marshal(order)
  if err != nil {
    log.Fatal(err)
  }

  // -> Context Header Injection
  headers := broker.InjectAMQPHeaders(amqpContext)

  // ->> Publish the results of the new Order to gRPC Channel:
  grpc.channel.PublishWithContext(
    amqpContext,
    "",
    q.Name,
    false,
    false,
    amqp.Publishing{
      ContentType  : "application/json",
      Body         : marshalledOrder,
      DeliveryMode : amqp.Persistent,
      Headers      : headers,
    },
  )

  return order, nil
}

func(grpc *grpcHandler) GetOrder(
  ctx context.Context, 
  req *pb.GetOrderRequest,
) (*pb.CreateOrderResponse, error) {
  return grpc.service.GetOrder(ctx, req)
}

func(grpc *grpcHandler) UpdateOrder(
  ctx context.Context,
  res *pb.CreateOrderResponse,
)( *pb.CreateOrderResponse, error ){
  return grpc.service.UpdateOrder(ctx, res)
}

func(grpc *grpcHandler) DeleteOrder(
  context.Context, 
  *pb.CreateOrderResponse,
)( error ) {

  return nil
}

