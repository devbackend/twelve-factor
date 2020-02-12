package main

import (
	"go.uber.org/zap"
	"os"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	port := os.Getenv("PORT")
	if port == "" {
		logger.Fatal("Empty port")
	}

	dsn := os.Getenv("DSN")
	if dsn == "" {
		logger.Fatal("Empty dsn")
	}
}
