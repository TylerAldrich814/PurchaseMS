package main

import (
	"fmt"
	"net/http"
  "github.com/TylerAldrich814/common"
	pb "github.com/TylerAldrich814/common/api"
	"github.com/TylerAldrich814/gateway/gateway"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type handler struct {
  gateway gateway.OrdersGateway
}

func NewHandler(gateway gateway.OrdersGateway) *handler {
  return &handler{gateway}
}

func(h *handler) registerRoutes(mux *http.ServeMux){
  // ->> Serving Static Files
  mux.Handle("/", http.FileServer(http.Dir("public")))

  mux.HandleFunc(
    "POST /api/customers/{customerID}/orders", 
    h.handleCreateOrder,
  )
  mux.HandleFunc(
    "GET /api/customers/{customerID}/orders/{orderID}",
    h.handleGetOrder,
  )
}
func(h *handler) handleGetOrder(w http.ResponseWriter, r *http.Request) {
  customerID := r.PathValue("customerID")
  orderID := r.PathValue("orderID")

  order, err := h.gateway.GetOrder(r.Context(), customerID, orderID)
  if order == nil {
    common.WriteError(w,
      http.StatusNotFound, 
      fmt.Sprintf("Customer %s Failed to Get OrderID %s", customerID, orderID),
    )
    return
  }
  grpcStatus := status.Convert(err)
  if grpcStatus != nil {
    if grpcStatus.Code() != codes.InvalidArgument {
      common.WriteError(w, http.StatusBadRequest, grpcStatus.Message())
    }
    common.WriteError(w, http.StatusInternalServerError, err.Error())
    return
  }
  common.WriteJSON(w, http.StatusOK, order)
  return
}

func(h *handler) handleCreateOrder(w http.ResponseWriter, r *http.Request) {
  customerID := r.PathValue("customerID")

  var items []*pb.ItemsWithQuantity
  if err := common.ReadJSON(r, &items); err != nil {
    common.WriteError(w, http.StatusBadRequest, err.Error())
    return
  }

  if err := validateItems(items); err != nil {
    common.WriteError(w, http.StatusBadRequest, err.Error())
    return 
  }

  res, err := h.gateway.CreateOrder(
    r.Context(), 
    &pb.CreateOrderRequest{
      CustomerId: customerID,
      Items: items,
    },
  )

  grpcStatus := status.Convert(err)
  if grpcStatus != nil {

    if grpcStatus.Code() != codes.InvalidArgument {
      common.WriteError(w, http.StatusBadRequest, grpcStatus.Message())
      return
    }
    common.WriteError(w, http.StatusInternalServerError, err.Error())
    return
  }
  orderRes := &CreateOrderRequest{
    Order: res,
    RedirectURL: fmt.Sprintf(
      "http://localhost:8080/success.html?customerID=%s&orderID=%s",
      res.CustomerId,
      res.Id,
    ),
  }

  common.WriteJSON(w, http.StatusOK, orderRes)
}

func validateItems(items []*pb.ItemsWithQuantity) error {
  if len(items) == 0 {
    return common.ErrorNoItems
  }

  for _, i := range items {
    if i.Id == "" {
      return common.ErrorIdRequired
    }
    if i.Quantity <= 0 {
      return common.ErrorQuantityBelowOne
    }
  }

  return nil
}
