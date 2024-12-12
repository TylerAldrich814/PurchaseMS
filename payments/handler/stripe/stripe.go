package handler

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"

  pb "github.com/TylerAldrich814/common/api"
	"github.com/TylerAldrich814/common"
	"github.com/TylerAldrich814/common/broker"
	_ "github.com/joho/godotenv/autoload"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/webhook"
)

var (
  stripeMode     = common.EnvString("STRIPE_MODE", "DEVELOPMENT")
  stripeSecret = common.EnvString("STRIPE_SECRET", "")
)

type StripePaymentHandler struct {
  secret  string
  execCmd *exec.Cmd
  channel *amqp.Channel
}

func NewStripePaymentHandler(
  ctx     context.Context,
  url     string,
  channel *amqp.Channel,
) *StripePaymentHandler {

  var secret string
  log.Println(stripeSecret)
  if stripeSecret == "" {
    log.Printf("Obtaining Stripe Secret..")
    s, err := extractStripeSecret(url)
    if err != nil {
      panic(err)
    }
    log.Printf("Obtained Secret: \"%s\"", s)
    secret = s
  } else {
    secret = stripeSecret
  }

  var execCmd *exec.Cmd
  if stripeMode == "DEVELOPMENT" {
    execCmd = exec.Command(
      "stripe", 
      "listen", 
      "--forward-to", 
      url,
    )
    startStripeForwardPort(ctx, url, execCmd)
  }

  return &StripePaymentHandler{ 
    secret, 
    execCmd,
    channel,
  }
}

func(s *StripePaymentHandler) AwaitForShutdown() {
  if s.execCmd == nil { return }
  go func(){
    // ->> Channel to listen for system interrupt signals 
    sigs := make(chan os.Signal, 1)
    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

    // ->> Await for interrupt signal to gracefully shutdown
    <-sigs

    // Cleanup: Stop the Stripe CLI process when shutting down the application
    if s.execCmd != nil {
      if err := s.execCmd.Process.Kill(); err != nil {
        log.Printf("Error stopping Stripe CLI: %v", err)
      } else {
        log.Println("Stripe CLI process stopped successfully")
      }
    }
  }()
}

func(s *StripePaymentHandler) HandleCheckout(
  w http.ResponseWriter, 
  r *http.Request,
) {
  const MaxBosyBytes = int64(65536)
  r.Body = http.MaxBytesReader(w, r.Body, MaxBosyBytes)

  body, err := io.ReadAll(r.Body)
  if err != nil {
    log.Printf("Stripe::HandleCheckout: Error reading Request Body: %s", err.Error())
    common.WriteError(w, http.StatusBadRequest, "Failed to read Request Body")
    return
  }
  log.Printf("Stripe::HandleCheckout: Successful recieved Body: %s\n", body)

  event, err := webhook.ConstructEvent(
    body, 
    r.Header.Get("Stripe-Signature"),
    s.secret,
  )
  if err != nil {
    log.Printf("Stripe::HandleCheckout: Error Verifying Stripe Secret: %s", err.Error())
    common.WriteError(w, http.StatusBadRequest, "Failed to verify Stripe Secret")
    return
  }
  log.Printf("Event: %s", event.Type)

  if event.Type == stripe.EventTypeCheckoutSessionExpired {
    log.Printf("Stripe::HandleCheckout: Error Stripe Secret Verification Expired")
    common.WriteError(w, http.StatusInternalServerError, "Stripe Secret Verification Expired")
    return
  } else if event.Type == stripe.EventTypeCheckoutSessionCompleted {
    var session stripe.CheckoutSession
    if err := json.Unmarshal(event.Data.Raw, &session); err != nil {
      log.Printf("Stripe::HandleCheckout: Error Parsing Webhook JSON - %s", err.Error())
      common.WriteError(w, http.StatusBadRequest, "Failed to Parse Strupe Webhook JSON")
      return
    }
    if session.PaymentStatus == "paid" {
      log.Printf("Payment for Checkout Session %v succeeded.", session.ID)

      orderID := session.Metadata["orderID"]
      customerID := session.Metadata["customerID"]

      ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
      defer cancel()

      order := &pb.CreateOrderResponse{
        Id          : orderID,
        CustomerId  : customerID,
        Status      : "paid",
        PaymentLink : "",
      }

      marshalledOrder, err := json.Marshal(order)
      if err != nil {
        log.Printf("Stripe::handleCheckout - Failed to Marshal Order Response - %s", err.Error())
        common.WriteError(w, http.StatusInternalServerError, "Failed to Marshal Order Response")
        return
      }

      s.channel.PublishWithContext(
        ctx,
        broker.OrderPaidEvent,
        "",
        false,
        false,
        amqp.Publishing{
          ContentType  : "application/json",
          Body         : marshalledOrder,
          DeliveryMode : amqp.Persistent,
        },
      )
    }
    log.Printf("Stripe::handleCheckout - Message published \"order.paid\"")
  }

  w.WriteHeader(http.StatusOK)
}
