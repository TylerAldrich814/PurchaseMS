package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"time"

	"github.com/TylerAldrich814/common"
	"github.com/TylerAldrich814/common/broker"
	"github.com/TylerAldrich814/common/discovery"
	"github.com/TylerAldrich814/common/discovery/consul"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var (
  serviceName = "stock"
  grpcAddr    = common.EnvString("GRPC_ADDR",     "localhost:2022")
  consulAddr  = common.EnvString("CONSUL_ADDR",   "localhost:8500")
  jaegerAddr  = common.EnvString("JAEGER_ADDR",   "localhost:4318")
  amqpUser    = common.EnvString("RABBITMQ_USER", "guest")
  amqpPass    = common.EnvString("RABBITMQ_PASS", "guest")
  amqpHost    = common.EnvString("RABBITMQ_HOST", "localhost")
  amqpPort    = common.EnvString("RABBITMQ_PORT", "5672")
)

func main(){
  ctx, cancel := signal.NotifyContext(
    context.Background(),
    os.Interrupt,
  )
  defer cancel()

  logger, _ := zap.NewProduction()
  defer logger.Sync()

  if err := common.SetGlobalTracer(
    ctx,
    serviceName,
    jaegerAddr,
  ); err != nil {
    logger.Fatal("Could not set Global Tracker", zap.Error(err))
  }

  registry, err  := consul.NewRegistry(consulAddr, serviceName)
  if err != nil {
    panic(err)
  }

  instanceID := discovery.GenerateInstanceID(serviceName)
  if err := registry.Register(
    ctx, 
    instanceID, 
    serviceName, 
    grpcAddr,
  ); err != nil {
    panic(err)
  }

  // ->> Registry Health Check
  go func(){
    for {
      if err := registry.HealthCheck(instanceID, serviceName); err != nil {
        logger.Error("Failed to run Health Check", zap.Error(err))
      }
      time.Sleep(time.Second * 2)
    }
  }()
  defer registry.Deregister(ctx, instanceID, serviceName)

  channel, close := broker.Connect(amqpUser, amqpPass, amqpHost, amqpPort)
  defer func(){
    close()
    channel.Close()
  }()

  grpcServer := grpc.NewServer()

  listener, err := net.Listen("tcp", grpcAddr)
  if err != nil {
    logger.Fatal(
      fmt.Sprintf("TCP: Failed to Listen on addreess %s", grpcAddr), 
      zap.Error(err),
    )
  }
  defer listener.Close()

  store := NewStore()
  svc := NewService(store)
  svcWithTelemetry := NewTelemetryMiddleware(svc)
  svcWithLogging   := NewLoggingMiddleware(svcWithTelemetry)

  NewGRPCHandler(grpcServer, channel, svcWithLogging)

  consumer := NewConsumer()
  go consumer.Listen(channel)

  logger.Info("Starting gRPC Server @", zap.String("port", grpcAddr))

  if err := grpcServer.Serve(listener); err != nil {
    logger.Fatal("Failed to server", zap.Error(err))
  }
}
