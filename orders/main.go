package main

import (
	"context"
	"net"
	"os"
	"os/signal"
	"time"

	"github.com/TylerAldrich814/common"
	"github.com/TylerAldrich814/common/broker"
	"github.com/TylerAldrich814/common/discovery"
	"github.com/TylerAldrich814/common/discovery/consul"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
  "go.uber.org/zap"
)

var (
  serviceName = "orders"
  grpcAddr    = common.EnvString("GRPC_ADDR", "localhost:2000")
  consulAddr  = common.EnvString("CONSUL_ADDR", "localhost:8500")
  jaegerAddr  = common.EnvString("JAEGAR_ADDR", "localhost:4318")
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

  // ->> Zap Logger Setup:
  logger, _ := zap.NewProduction()
  defer logger.Sync()
  zap.ReplaceGlobals(logger)

  // ->> Global Jaeger Telemetry:
  if err := common.SetGlobalTracer(ctx, serviceName, jaegerAddr); err != nil {
    logger.Fatal("Failed to set Global Tracer: %v", zap.Error(err))
  }

  // ->> Consul Network Mesh Registration:
  registry, err := consul.NewRegistry(
    consulAddr,
    serviceName,
  )
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

  go func(){
    for {
      if err := registry.HealthCheck(instanceID, serviceName); err != nil {
        logger.Fatal("Failed Healt Check")
      }
      time.Sleep(time.Second * 2)
    }
  }()
  defer registry.Deregister(ctx, instanceID, serviceName)

  // ->> RabbitMQ Broker Connection:
  channel, close := broker.Connect(
    amqpUser,
    amqpPass,
    amqpHost,
    amqpPort,
  )
  defer func(){
    close()
    channel.Close()
  }()

  grpcServer := grpc.NewServer()

  listener, err := net.Listen("tcp", grpcAddr)
  if err != nil {
    logger.Fatal("Failed to listen: %v\n", zap.Error(err))
  }
  defer listener.Close()

  store := NewStore()
  svc := NewService(store)
  svcWithTelemetry := NewTelemetryMiddleWare(svc)
  svcWithLogging := NewLoggerMiddleware(svcWithTelemetry)

  NewGRPCHandler(grpcServer, svcWithLogging, channel)

  // ->> amqp Consumer Connection:
  amqpConsumer := NewConsumer(svcWithLogging)
  go amqpConsumer.Listen(channel)

  logger.Info("->> Starting Orders Service @ ", zap.String("port",grpcAddr))
  if err := grpcServer.Serve(listener); err != nil {
    logger.Fatal(err.Error())
  }
}
