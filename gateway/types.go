package main

import (
  pb "github.com/TylerAldrich814/common/api"
)

type CreateOrderRequest struct {
  Order       *pb.CreateOrderResponse `json:"order"`
  RedirectURL string                  `json:"redirect_url"`
}
