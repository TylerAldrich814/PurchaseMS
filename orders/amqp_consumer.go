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
  service OrdersService
}

func NewConsumer(service OrdersService) *consumer {
  return &consumer{ service }
}

func(c *consumer) Listen(channel *amqp.Channel) {
  q, err := channel.QueueDeclare("", true, false, true, false, nil)
  if err != nil {
    log.Fatal(err)
  }

  if err := channel.QueueBind(
    q.Name, 
    "", 
    broker.OrderPaidEvent, 
    false, 
    nil,
  ); err != nil {
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
    log.Print("LISTENING AMQP CONSUMER")
    for msg := range msgs {
      log.Printf("Received Message: %s", msg.Body)

      // -> Extract Headers
      ctx := broker.ExtractAMQPHeaders(context.Background(), msg.Headers)
      tr := otel.Tracer("amqp")
      _, span := tr.Start(
        ctx,
        fmt.Sprintf(
          "AMQP consume - %s",
          q.Name,
        ),
      )

      order := &pb.CreateOrderResponse{}
      if err := json.Unmarshal(msg.Body, order); err != nil {
        msg.Nack(false, false)
        log.Printf("Failed to Unmarshal Order: %s", err.Error())
        continue
      }

      _, err := c.service.UpdateOrder(context.Background(), order)
      if err != nil {
        log.Println("Failed to update order: %s", err.Error())
        if err := broker.HandleRetry(channel, &msg); err != nil {
          log.Printf("Error Handling Retry: %v", err)
        }
        continue
      }

      span.AddEvent("order.updated")
      span.End()

      log.Println("Order has been updated from AMQP")
      msg.Ack(false)
    }
  }()
  <-forever
}
