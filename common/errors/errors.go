package errors

import "errors"

var (
  NoItems              = errors.New("Items must have at least one item")
  IdRequired           = errors.New("ID is Required")
  QuantityBelowOne     = errors.New("Quantity must be above 0")
  OrderNotFound        = errors.New("Requested Order was not found in Storage")
  AMQP_RetryFailed     = errors.New("RabbitMQ Message: Ran out of retries")
  Stock_ItemNotFound   = errors.New("Stock: Item Not Found")
  Stock_ItemOutOfStock = errors.New("Stock: Item Not in Stock")
)
