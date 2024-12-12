package main

import (
	"context"
	"testing"

	"github.com/TylerAldrich814/common/api"
	testInmem "github.com/TylerAldrich814/common/broker/inmem"
	"github.com/TylerAldrich814/payments/gateway"
	"github.com/TylerAldrich814/payments/processor/inmem"
)


func TestService(t *testing.T) {
  inMemPaymentProcessor := inmem.NewInMemory()
  registry := testInmem.NewRegistry()
  gateway  := gateway.NewGateway(registry)
  svc      := NewService(inMemPaymentProcessor, gateway)

  t.Run("Should Create a Payment Link", func(t *testing.T){
    link, err := svc.CreatePayments(context.Background(), &api.CreateOrderResponse{
      Id: "10",
      CustomerId: "1",
    })
    if err != nil {
      t.Errorf("CreatePayment() error = %v, want nil", err)
    }
    if link == "" {
      t.Errorf("CreatePayment() Link is Empty")
    }
  })
}
