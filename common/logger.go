package common

import (
  "time"
	"go.uber.org/zap"
)

func LogInfo(
  msg string,
) func() {
  start := time.Now()

  return func(){
    zap.L().Info(msg, zap.Duration("took", time.Since(start)))
  }
}
