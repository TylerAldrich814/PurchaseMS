package broker

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/TylerAldrich814/common/errors"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.opentelemetry.io/otel"
)

const (
  Queue         = "main_queue"
  DLQ           = "dlq_main"
  DLX           = "dlx_main"
  recountKey    = "x-retry-count"
  maxRetryCount = 3
)

func Connect(user, pass, host, port string)(  *amqp.Channel, func()error ){
  address := fmt.Sprintf(
    "amqp://%s:%s@%s:%s",
    user,
    pass,
    host,
    port,
  )
  connection, err := amqp.Dial(address)
  if err != nil {
    log.Fatal(err)
  }

  channel, err := connection.Channel()
  if err != nil {
    log.Fatal(err)
  }

  if err := channel.ExchangeDeclare(
    OrderCreatedEvent, 
    "direct", 
    true,
    false,
    false,
    false,
    nil,
  ); err != nil {
    log.Fatal(err)
  }

  if err := channel.ExchangeDeclare(
    OrderPaidEvent, 
    "fanout", 
    true,
    false,
    false,
    false,
    nil,
  ); err != nil {
    log.Fatal(err)
  }

  if err := createDLQAndDLX(channel); err != nil {
    log.Fatal("Failed to create DLQ/DLX: %v", err)
  }

  return channel, connection.Close
}

func HandleRetry(
  channel  *amqp.Channel,
  delivery *amqp.Delivery,
) error {
  // ->> Creating/Updating the Retry Count:
  if delivery.Headers == nil {
    delivery.Headers = amqp.Table{}
  }

  retryCount, ok := delivery.Headers[recountKey].(int64)
  if !ok {
    retryCount = 0
  }
  retryCount++
  delivery.Headers[recountKey] = retryCount

  log.Printf("Retrying Message: Retry #%d", retryCount)

  if retryCount >= maxRetryCount {
    log.Printf("Moving Message to DLX: -- %s", delivery.Body)
    // DLQ
    channel.PublishWithContext(
      context.Background(),
      "",
      DLQ,
      false,
      false,
      amqp.Publishing{
        ContentType  : "application/json",
        Headers      : delivery.Headers,
        Body         : delivery.Body,
        DeliveryMode : amqp.Persistent,
      },
    )
    return errors.ErrorAMQP_RetryFailed
  }
  time.Sleep(time.Second * time.Duration(retryCount))

  return channel.PublishWithContext(
    context.Background(),
    delivery.Exchange,
    delivery.RoutingKey,
    false,
    false,
    amqp.Publishing{
      ContentType  : "application/json",
      Headers      : delivery.Headers,
      Body         : delivery.Body,
      DeliveryMode : amqp.Persistent,
    },
  )
}

// Creates both a Dead Letter Queue and a Dead Letter Exchange
// For Messages that failed *maxRetryCount* times during HandleRetry 
//
// ...
func createDLQAndDLX(channel *amqp.Channel) error {
  // ->> Create DLQ Main Queue
  queue, err := channel.QueueDeclare(
    Queue,
    true,  // -> Durable
    false, // -> Delete When Used
    false, // -> Exclusive
    false, // -> No-Wait
    nil,   // -> Arguments
  )
  if err != nil {
    return err
  }

  // ->> Declare Dead Letter Exchange
  if err := channel.ExchangeDeclare(
    DLX,
    "fanout", // -> DX Type
    true,     // -> Durable
    false,    // -> Auto-Deleted
    false,    // -> Internal
    false,    // -> No-Wait
    nil,      // -> Arguments
  ); err != nil {
    return err
  }

  // ->> Bind Main Queue and DLX
  if err := channel.QueueBind(
    queue.Name,
    "",
    DLX,
    false,
    nil,
  ); err != nil {
    return err
  }

  // ->> Declare Dead Letter Queue
  if _, err := channel.QueueDeclare(
    DLQ,
    true,  // -> Durable
    false, // -> Delete-When-Unused
    false, // -> Exclusive
    false, // -> No-Wait
    nil,   // -> Arguments
  ); err != nil {
    return err
  }

  return nil
}

type AmqpHeaderCarrier map[string]interface{}

func(a AmqpHeaderCarrier) Get(key string) string {
  value, ok := a[key]
  if !ok {
    return ""
  }

  return value.(string)
}

func(a AmqpHeaderCarrier) Set(key, val string) {
  a[key] = val
}

func(a AmqpHeaderCarrier) Keys() []string {
  keys := make([]string, len(a))
  i := 0
  for key := range a {
    keys[i] = key
    i++
  }

  return keys
}

func InjectAMQPHeaders(
  ctx context.Context,
) map[string]interface{} {
  carrier := AmqpHeaderCarrier{}
  otel.GetTextMapPropagator().Inject(ctx, carrier)

  return carrier
}

func ExtractAMQPHeaders(
  ctx     context.Context,
  headers map[string]interface{},
) context.Context {
  return otel.GetTextMapPropagator().Extract(ctx, AmqpHeaderCarrier(headers))
}
