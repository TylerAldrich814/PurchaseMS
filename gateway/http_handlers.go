package main

import (
	"net/http"

	"github.com/TylerAldrich814/common"
	pb "github.com/TylerAldrich814/common/api"
	"github.com/TylerAldrich814/gateway/gateway"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type handler struct {
  // gateway
  gateway gateway.OrdersGateway
}

func NewHandler(gateway gateway.OrdersGateway) *handler {
  return &handler{gateway}
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

  // req := &CreateOrderRequest{
  //   Order : res,
  //   // RedirectURL: fmt.Sprintf("")
  // }

  common.WriteJSON(w, http.StatusOK, res)
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
