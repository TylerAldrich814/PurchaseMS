package main

import (
	"context"
	"log"
	"net"

	"github.com/TylerAldrich814/common"
	"google.golang.org/grpc"
)

var (
  grpcAddr = common.EnvString("GRPC_ADDR", "localhost:2000")
)

func main(){
  grpcServer := grpc.NewServer()

  listener, err := net.Listen("tcp", grpcAddr)
  if err != nil {
    log.Fatal("Failed to listen: %s", grpcAddr)
  }
  defer listener.Close()

  store := NewStore()
  svc := NewService(store)
  NewGRPCHandler(grpcServer)

  svc.CreateOrder(context.Background())

  log.Printf("Orders:: Server is Running @ %s\n", grpcAddr)
  if err := grpcServer.Serve(listener); err != nil {
    log.Fatal(err.Error())
  }
}
