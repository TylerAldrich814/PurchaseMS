package main

import (
	"net/http"

	"github.com/TylerAldrich814/payments/handler"
	amqp "github.com/rabbitmq/amqp091-go"
)

type PaymentHTTPHandler struct {
  channel *amqp.Channel
  handler handler.PaymentHandler
}

func NewPaymentHTTPHandler(
  channel *amqp.Channel,
  handler handler.PaymentHandler,
) *PaymentHTTPHandler {
  return &PaymentHTTPHandler{ 
    channel,
    handler,
  }
}

func(h *PaymentHTTPHandler) registerRoutes(router *http.ServeMux){
  router.HandleFunc("/webhook", h.handleCheckoutWebhook)
}

func(h *PaymentHTTPHandler) handleCheckoutWebhook(
  w http.ResponseWriter, 
  r *http.Request,
){
  h.handler.HandleCheckout(w, r)
}
