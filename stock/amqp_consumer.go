package main

import (
	"context"
	"fmt"
	"log"

	"github.com/TylerAldrich814/common/broker"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.opentelemetry.io/otel"
)

type Consumer struct {}

func NewConsumer() *Consumer {
  return &Consumer{}
}

func(c *Consumer) Listen(channel *amqp.Channel) {
  q, err := channel.QueueDeclare(
    "",    // -> Name
    true,  // -> Durable
    false, // -> Delete-When-Unused
    true,  // -> Exclusive
    false, // -> No-Wait
    nil,   // -> Arguments
  )
  if err != nil {
    log.Fatal(err)
  }

  if err := channel.QueueBind(
    q.Name,                 // -> Queue Name
    "",                     // -> Routing Key
    broker.OrderPaidEvent,  // -> Exchange
    false,                  // -> No-Wait
    nil,                    // -> Queue Table
  ); err != nil {
    log.Fatal(err)
  }

  msgs, err := channel.Consume(
    q.Name,  // -> Queue Name
    "",      // -> Consumer Name
    false,   // -> Auto Acknowledge
    false,   // -> Exclusive
    false,   // -> No-Local
    false,   // -> No-Wait
    nil,     // -> Table Arguments
  )
  if err != nil {
    log.Fatal(err)
  }

  var forever chan struct{}
  go func(){
    for msg := range msgs {
      ctx := broker.ExtractAMQPHeaders(context.Background(), msg.Headers)

      tr := otel.Tracer("amqp")
      _, messageSpan := tr.Start(
        ctx,
        fmt.Sprintf("AMQP - consume - %s", q.Name),
      )
      log.Printf("Received Message: %s", msg.Body)
      orderID := string(msg.Body)
      msg.Ack(false)

      messageSpan.End()
      log.Printf("Order Received: %s", orderID)
    }
  }()
  log.Printf("AMQP Listening: Ctrl+C to Exit")
  <-forever
}
