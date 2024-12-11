package common

import "errors"

var (
  ErrorNoItems          = errors.New("Items must have at least one item")
  ErrorIdRequired       = errors.New("ID is Required")
  ErrorQuantityBelowOne = errors.New("Quantity must be above 0")
  ErrorOrderNotFound    = errors.New("Requested Order was not found in Storage")
)
