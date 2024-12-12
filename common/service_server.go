package common

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"

	"github.com/TylerAldrich814/common/broker"
	"github.com/TylerAldrich814/common/discovery"
	"github.com/TylerAldrich814/common/discovery/consul"
)

var (
  grpcAddr    = EnvString("GRPC_ADDR", "localhost:2000")
  consulAddr  = EnvString("CONSUL_ADDR", "localhost:8500")
  amqpUser    = EnvString("RABBITMQ_USER", "guest")
  amqpPass    = EnvString("RABBITMQ_PASS", "guest")
  amqpHost    = EnvString("RABBITMQ_HOST", "localhost")
  amqpPort    = EnvString("RABBITMQ_PORT", "5672")
)

type ServiceServiceConfig struct {
  Address    string
  MainHost   string
  BrokerUser string
  BrokerPass string
  BrokerHost string
  BrokerPort string
}


type ServiceServer struct {
  serviceName      string
  instanceID       string
  ctx              context.Context
  ctxCancel        context.CancelFunc
  listener         net.Listener
  registry         *consul.Registry
  channel          *amqp.Channel
  channelClose     func() error
  grpcServer       *grpc.Server
}

func(s *ServiceServer) Serve(
  handlerSetup func(grpcServer *grpc.Server, channel *amqp.Channel),
) error {
  // Health Check
  go func(){
    for {
      if err := s.registry.HealthCheck(
        s.instanceID,
        s.serviceName,
      ); err != nil {
        log.Fatalf("%s Health Check Failed: %v", s.serviceName, err)
      }
      time.Sleep(time.Second * 2)
    }
  }()

  handlerSetup(s.grpcServer, s.channel)
  return s.grpcServer.Serve(s.listener)
}

func(s *ServiceServer) Stop() {
  s.registry.Deregister(s.ctx, s.instanceID, s.serviceName)
  s.channelClose()
  s.listener.Close()
}

func NewServiceServer(
  serviceName string,
  config      *ServiceServiceConfig,
)( *ServiceServer, error ){
  // Congiguration

  if config == nil {
    config = &ServiceServiceConfig{
      Address    : grpcAddr,
      MainHost   : consulAddr,
      BrokerUser : amqpUser,
      BrokerPass : amqpPass,
      BrokerHost : amqpHost,
      BrokerPort : amqpPort,
    }
  } else {
    if config.Address == "" {
      config.Address = grpcAddr
    }
    if config.MainHost == "" {
      config.MainHost = consulAddr
    }
    if config.BrokerUser == "" || config.BrokerPass == "" {
      config.BrokerHost = amqpUser
      config.BrokerPass = amqpPass
    }
    if config.BrokerHost == "" {
      config.BrokerHost = amqpHost
    }
    if config.BrokerPort == "" {
      config.BrokerPort = amqpPort
    }
  }  

  ctx, ctxCancel := signal.NotifyContext(
    context.Background(),
    os.Interrupt,
  )
  registry, err := consul.NewRegistry(
    consulAddr,
    serviceName,
  )
  if err != nil {
    return nil, fmt.Errorf("Failed to create Registry for %s - %w", serviceName, err)
  }
  
  instanceID := discovery.GenerateInstanceID(serviceName)

  if err := registry.Register(
    ctx,
    instanceID,
    serviceName,
    config.Address,
  ); err != nil {
    return nil, fmt.Errorf("Failed to Register %s: %w", serviceName, err)
  }

  channel, channelClose := broker.Connect(
    config.BrokerUser,
    config.BrokerPass,
    config.BrokerHost,
    config.BrokerPort,
  )

  grpcServer := grpc.NewServer()

  listener, err := net.Listen("tcp", config.Address)
  if err != nil {
    return nil, fmt.Errorf("Failed start TCP for %s - %w", serviceName, err)
  }

  return &ServiceServer{
    serviceName,
    instanceID,
    ctx,
    ctxCancel,
    listener,
    registry,
    channel,
    channelClose,
    grpcServer,
  }, nil
}
