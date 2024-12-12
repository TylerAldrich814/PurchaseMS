package main

import (
	"context"
	"encoding/json"
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
  // ->> Validate incoming Order(s)
  items, err := grpc.service.ValidateOrder(ctx, req)
  if err != nil {
    return nil, err
  }
  
  // ->> Store Incoming Order(s) in Order Service DB
  order, err := grpc.service.CreateOrder(ctx, req, items)
  if err != nil {
    return nil, err
  }

  marshalledOrder, err := json.Marshal(order)
  if err != nil {
    log.Fatal(err)
  }

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

  // ->> Publish the results of the new Order to gRPC Channel 
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




