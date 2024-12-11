package handler

import "net/http"

type PaymentHandler interface {
  HandleCheckout(w http.ResponseWriter, r *http.Request)
}
