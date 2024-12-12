package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/TylerAldrich814/common"
	"github.com/TylerAldrich814/common/broker"
	"github.com/TylerAldrich814/common/discovery"
	"github.com/TylerAldrich814/common/discovery/consul"
	"github.com/TylerAldrich814/payments/gateway"
	stripeHandler "github.com/TylerAldrich814/payments/handler/stripe"
	stripeProcessor "github.com/TylerAldrich814/payments/processor/stripe"
	_ "github.com/joho/godotenv/autoload"
	"github.com/stripe/stripe-go/v81"
	"google.golang.org/grpc"
)

func Env(key, fallback string) string {
  if val, ok := syscall.Getenv(key); ok {
    return val
  }

  return fallback
}

var (
  serviceName    = "payments"
  stripeKey      = common.EnvString("STRIPE_KEY", "")
  httpAddr       = common.EnvString("HTTP_ADDR",     "localhost:8081")
  grpcAddr       = common.EnvString("GRPC_ADDR",     "localhost:2001")
  consulAddr     = common.EnvString("CONSUL_ADDR",   "localhost:8500")
  amqpUser       = common.EnvString("RABBITMQ_USER", "guest")
  amqpPass       = common.EnvString("RABBITMQ_PASS", "guest")
  amqpHost       = common.EnvString("RABBITMQ_HOST", "localhost")
  amqpPort       = common.EnvString("RABBITMQ_PORT", "5672")
)


func main(){
  ctx, cancel := signal.NotifyContext(
    context.Background(),
    os.Interrupt,
  )
  defer cancel()

  if stripeKey == "" {
    panic("NO STRIPE KEY FOUND")
  }

  // ->> Consul Network Mesh Registration::
  registry, err := consul.NewRegistry(consulAddr, serviceName)
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
        log.Fatal("Failed Health Check")
      }
      time.Sleep(time.Second*2)
    }
  }()
  defer registry.Deregister(ctx, instanceID, serviceName)

  // ->> Stripe Connection::
  stripe.Key = stripeKey

  // ->> RabbitMQ Broker Connection::
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

  // ->> Payment Service Connection::
  stripeProcessor := stripeProcessor.NewProcessor()
  gateway := gateway.NewGateway(registry)
  svc := NewService(stripeProcessor, gateway)

  // ->> RabbitMQ Consumer Connection::
  amqpConsumer := NewConsumer(svc)
  go amqpConsumer.Listen(channel)

  // ->> Payment HTTP Handler Server
  mux := http.NewServeMux()

  // ->> PaymentHandler Connection:
  stripePaymentHandler := stripeHandler.
    NewStripePaymentHandler(
      ctx,
      fmt.Sprintf("%s/webhook", httpAddr),
      channel,
    )
  stripePaymentHandler.AwaitForShutdown()

  paymentHandler := NewPaymentHTTPHandler(
    channel,
    stripePaymentHandler,
  )
  paymentHandler.registerRoutes(mux)

  go func(){
    log.Printf("->> Started HTTP Server @ %s", httpAddr)
    if err := http.ListenAndServe(httpAddr, mux); err != nil {
      log.Fatalf("Failed to start HTTP Server: %v", err)
    }
  }()

  // ->> gRPC Server Connection::
  grpcServer := grpc.NewServer()

  listener, err := net.Listen("tcp", grpcAddr)
  if err != nil {
    log.Fatalf("Failed to Listen: %v\n", err)
  }
  defer listener.Close()

  log.Printf("->> Starting Payments Service @ %s..\n", grpcAddr)
  if err := grpcServer.Serve(listener); err != nil {
    log.Fatal(err.Error())
  }
}
