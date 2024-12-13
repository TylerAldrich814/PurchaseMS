package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	pb "github.com/TylerAldrich814/common/api"
	"github.com/TylerAldrich814/common/broker"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.opentelemetry.io/otel"
)

type consumer struct {
  service PaymentService
}

func NewConsumer(service PaymentService) *consumer {
  return &consumer{ service }
}

func(c *consumer) Listen(channel *amqp.Channel) {

  q, err := channel.QueueDeclare(
    broker.OrderCreatedEvent,
    true,
    false,
    false,
    false,
    nil,
  )
  if err != nil {
    log.Fatal(err)
  }

  msgs, err := channel.Consume(
    q.Name, 
    "", 
    false,
    false,
    false,
    false,
    nil,
  )
  if err != nil {
    log.Fatal(err)
  }

  var forever chan struct{}

  go func(){
    for msg := range msgs {
      // d.Ack()
      log.Printf("Received Message: %v", msg)

      // -> Extract Headers
      ctx := broker.ExtractAMQPHeaders(context.Background(), msg.Headers)
      tr := otel.Tracer("amqp")
      _, msgSpan := tr.Start(
        ctx,
        fmt.Sprintf(
          "AMQP - consume - %s",
          q.Name,
        ),
      )

      order := &pb.CreateOrderResponse{}
      if err := json.Unmarshal(msg.Body, &order); err != nil {
        msg.Nack(false, false)
        log.Printf("Failed to unmarchal order: %v", err)
        continue;
      }

      paymentLink, err := c.service.CreatePayments(
        context.Background(),
        order,
      )
      if err != nil {
        log.Printf("Failed to create payment: %v", err)

        if err := broker.HandleRetry(channel, &msg); err != nil {
          log.Printf("Error Handling Retry: %v", err)
        }

        msg.Nack(false, false)
        continue
      }

      msgSpan.AddEvent(fmt.Sprintf(
        "payment.created: %s",
        paymentLink,
      ))
      msgSpan.End()

      log.Printf("Payment Link Created:  %s", paymentLink)
      msg.Ack(false)
    }
  }()
  <-forever
}
