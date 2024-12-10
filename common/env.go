package common

import (
  "syscall"
  "log"
)

func EnvString(key, fallback string) string {
  if val, ok := syscall.Getenv(key); ok {
    return val
  }

  log.Printf("FAILED to find env \"%s\"", key)
  return fallback
}
