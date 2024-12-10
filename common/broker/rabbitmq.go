package broker

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
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

  if err := channel.ExchangeDeclare(OrderCreatedEvent, "direct", true, false, false, false, nil); err != nil {
    log.Fatal(err)
  }

  if err := channel.ExchangeDeclare(OrderPaidEvent, "direct", true, false, false, false, nil); err != nil {
    log.Fatal(err)
  }

  return channel, connection.Close
}
