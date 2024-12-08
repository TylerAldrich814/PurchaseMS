package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/TylerAldrich814/common"
	"github.com/TylerAldrich814/common/discovery"
	"github.com/TylerAldrich814/common/discovery/consul"
	"google.golang.org/grpc"
)

var (
  serviceName = "orders"
  grpcAddr    = common.EnvString("GRPC_ADDR", "localhost:2000")
  consulAddr  = common.EnvString("CONSUL_ADDR", "localhost:8500")
)

func main(){
  ctx, cancel := signal.NotifyContext(
    context.Background(),
    os.Interrupt,
  )
  defer cancel()

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
        log.Fatal("Failed Healt Check")
      }
      time.Sleep(time.Second * 2)
    }
  }()
  defer registry.Deregister(ctx, instanceID, serviceName)

  grpcServer := grpc.NewServer()

  listener, err := net.Listen("tcp", grpcAddr)
  if err != nil {
    log.Fatalf("Failed to listen: %s\n", grpcAddr)
  }
  defer listener.Close()

  store := NewStore()
  svc := NewService(store)
  NewGRPCHandler(grpcServer, svc)

  svc.CreateOrder(context.Background())

  log.Printf("Orders:: Server is Running @ %s\n", grpcAddr)
  if err := grpcServer.Serve(listener); err != nil {
    log.Fatal(err.Error())
  }
}
