package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	pb "github.com/TylerAldrich814/common/api"
	"github.com/TylerAldrich814/common/broker"
	amqp "github.com/rabbitmq/amqp091-go"
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
  log.Printf("New Order Received: %s", req)
  
  order, err := grpc.service.CreateOrder(ctx, req)

  marshalledOrder, err := json.Marshal(order)
  if err != nil {
    log.Fatal(err)
  }

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

  grpc.channel.PublishWithContext(
    ctx,
    "",
    q.Name,
    false,
    false,
    amqp.Publishing{
      ContentType  : "application/json",
      Body         : marshalledOrder,
      DeliveryMode : amqp.Persistent,
    },
  )

  return order, nil
}

func(grpc *grpcHandler) ValidateOrder(
  ctx   context.Context, 
  order *pb.CreateOrderRequest,
) error {
  if len(order.Items) == 0 {
    return errors.New("")
  }

  return nil
}
