package stripe

import (
	"fmt"
	"log"

	"github.com/TylerAldrich814/common"
	pb "github.com/TylerAldrich814/common/api"
	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/checkout/session"
)

var (
  gatewayHTTPAddr = common.EnvString("GATEWAY_HTTP_ADDRESS", "http://localhost:8080")
)

type Stripe struct{ }

func NewProcessor() *Stripe {
  return &Stripe{}
}

func(s *Stripe) CreatePaymentLink(
  res *pb.CreateOrderResponse,
)( string,error ){
  gatewaySuccessURL := fmt.Sprintf(
    "%s/success.html?customer_id=%s&order_id=%s", 
    gatewayHTTPAddr,
    res.CustomerId,
    res.Id,
  )
  gatewayCancelURL := fmt.Sprintf("%s/cancel.html", gatewayHTTPAddr)

  items := []*stripe.CheckoutSessionLineItemParams{}
  for _, item := range res.Items {
    log.Printf("-> ITEM: %s", item)
    items = append(items, &stripe.CheckoutSessionLineItemParams{
      Price    : stripe.String(item.PriceId),
      Quantity : stripe.Int64(int64(item.Quantity)),
    })
  }

  params := &stripe.CheckoutSessionParams{
    LineItems  : items,
    Mode       : stripe.String(string(stripe.CheckoutSessionModePayment)),
    SuccessURL : stripe.String(gatewaySuccessURL),
    CancelURL  : stripe.String(gatewayCancelURL),
  }

  result, err := session.New(params)
  if err != nil {
    log.Printf("Failed to create Payment Link: %v", err)
    log.Printf("RESULT: %s", result)
    return "", nil
  }
  return result.URL, nil
}
