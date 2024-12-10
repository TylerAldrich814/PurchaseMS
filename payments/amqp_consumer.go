package main

import (
	"context"
	"encoding/json"
	"log"

	pb "github.com/TylerAldrich814/common/api"
	"github.com/TylerAldrich814/common/broker"
	amqp "github.com/rabbitmq/amqp091-go"
)

type consumer struct {
  service PaymentService

}

func NewConsumer(service PaymentService) *consumer {
  return &consumer{ service }
}

func(c *consumer) Listen(ch *amqp.Channel) {
  q, err := ch.QueueDeclare(
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

  msgs, err := ch.Consume(
    q.Name, 
    "", 
    true, 
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

      order := &pb.CreateOrderResponse{}

      if err := json.Unmarshal(msg.Body, &order); err != nil {
        log.Printf("Failed to unmarchal order: %v\n", err)
        continue;
      }

      paymentLink, err := c.service.CreatePayments(
        context.Background(),
        order,
      )
      if err != nil {
        log.Printf("Failed to create payment: %v\n", err)
        continue
      }

      log.Printf("Payment Link Created:  %s\n", paymentLink)
    }
  }()
  <-forever
}
