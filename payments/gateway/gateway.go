package gateway

import "context"

type PaymentsGateway interface {
  UpdateOrderAfterPayment(ctx context.Context, orderID, paymentLink string) error
}
