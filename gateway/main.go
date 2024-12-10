package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/TylerAldrich814/common"
	"github.com/TylerAldrich814/gateway/gateway"
	_ "github.com/joho/godotenv/autoload"

	"github.com/TylerAldrich814/common/discovery"
	"github.com/TylerAldrich814/common/discovery/consul"
)

var (
  serviceName = "gateway"
  httpAddr    = common.EnvString("HTTP_ADDR", ":3000")
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
    httpAddr,
  ); err != nil {
    panic(err)
  }

  // ->> HealthCheck Routine
  go func(){
    for {
      if err := registry.HealthCheck(instanceID, serviceName); err != nil {
        log.Fatal("Failed Health Check")
      }
      time.Sleep(time.Second * 2)
    }
  }()
  defer registry.Deregister(ctx, instanceID, serviceName)

  mux := http.NewServeMux()
  gateway := gateway.NewGRPCGateway(registry)
  handler := NewHandler(gateway)
  handler.registerRoutes(mux)

  srv := &http.Server{
    Addr    : httpAddr,
    Handler : mux,
  }

  ch := make(chan error, 1)

  go func(){
    if err := srv.ListenAndServe(); err != nil {
      ch <- fmt.Errorf("Failed to Start Server: %w", err)
    }
  }()

  log.Printf("->> Starting Gateway @ %s..\n", httpAddr)
  select {
  case err := <-ch:
    log.Fatal(err)
  case <-ctx.Done():
    log.Println(" --> Shutting Down")
    timeout, cancel := context.WithTimeout(context.Background(), time.Second+10)
    defer cancel()

    if err := srv.Shutdown(timeout); err != nil {
      log.Fatalf("Failed to Shutdown Gateway Server: %w\n", err)
    } else {
      log.Println("Gracefully Shutdown Gateway Server ")
    }
  }
}
