package main

import (
	"net/http"

	"github.com/TylerAldrich814/common"
	pb "github.com/TylerAldrich814/common/api"
)

type handler struct {
  // gateway
  client pb.OrderServiceClient
}

func NewHandler(client pb.OrderServiceClient) *handler {
  return &handler{client}
}

func(h *handler) registerRoutes(mux *http.ServeMux){
  mux.HandleFunc("POST /api/customers/{customerID}/orders", h.HandleCreateOrder)
}

func(h *handler) HandleCreateOrder(w http.ResponseWriter, r *http.Request) {
  customerID := r.PathValue("customer_id")
  var items []*pb.ItemsWithQuantity
  if err := common.ReadJSON(r, &items); err != nil {
    common.WriteError(w, http.StatusBadRequest, err.Error())
    return
  }

  h.client.CreateOrder(r.Context(), &pb.CreateOrderRequest{
    CustomerId : customerID,
    Items      : items,
  })
}
