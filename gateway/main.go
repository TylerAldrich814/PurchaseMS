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
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

  pb "github.com/TylerAldrich814/common/api"
)

var (
  httpAddr = common.EnvString("HTTP_ADDR", ":2000")
  orderServiceAddr = "localhost:2000"
)

func main(){
  ctx, cancel := signal.NotifyContext(
    context.Background(),
    os.Interrupt,
  )
  defer cancel()

  conn, err := grpc.Dial(
    orderServiceAddr, 
    grpc.WithTransportCredentials(
      insecure.NewCredentials(),
    ),
  )
  if err != nil {
    log.Fatalf("Failed to dial gRPC Service at %s\n", orderServiceAddr)
  }
  defer conn.Close()

  log.Printf("Dialing to gRPC Service @ %s\n", orderServiceAddr)

  client := pb.NewOrderServiceClient(conn)

  mux := http.NewServeMux()
  handler := NewHandler(client)
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

  log.Printf("GATEWAY:: Starting HTTP Server @ %s..\n", httpAddr)

  select {
  case err := <-ch:
    log.Fatal(err)
  case <-ctx.Done():
    log.Println(" --> Shutting Down")
    timeout, cancel := context.WithTimeout(context.Background(), time.Second+10)
    defer cancel()

    if err := srv.Shutdown(timeout); err != nil {
      log.Fatal("Failed to Shutdown Gateway Server: %w", err)
    } else {
      log.Println("Gracefully Shutdown Gateway Server ")
    }
  }
}
// term := make(chan os.Signal, 1)
// signal.Notify(term, os.Interrupt, syscall.SIGTERM)
// go func(){
//   <-term
//   if err := srv.Close(); !errors.Is(err, http.ErrServerClosed){
//     log.Fatal("Error Closing Server: %w", err)
//   }
// }()
