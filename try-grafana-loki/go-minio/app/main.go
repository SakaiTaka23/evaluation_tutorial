package main

import (
	"math/rand"
	"time"

	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()

	for {
		switch i := rand.Intn(4); i {
		case 0:
			logger.Debug("Debug Level Message", zap.String("file", "main.go"))
		case 1:
			logger.Info("Info Level Message", zap.String("file", "main.go"))
		case 2:
			logger.Warn("Warn Level Message", zap.String("file", "main.go"))
		case 3:
			logger.Error("Error Level Message", zap.String("file", "main.go"))
		}

		time.Sleep(5 * time.Second)
	}
}
