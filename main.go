package main

import (
	"context"

	"github.com/FreitasGabriel/anotai-test-consumer/src/configuration/database"
	"github.com/FreitasGabriel/anotai-test-consumer/src/configuration/logger"
	"github.com/FreitasGabriel/anotai-test-consumer/src/service/queue"
	"github.com/joho/godotenv"
)

func main() {

	logger.Info("Starting server...")
	err := godotenv.Load()
	if err != nil {
		logger.Error(
			"error to initialize environment variables",
			err,
		)
		return
	}

	db, err := database.InitDatabase(context.Background())
	if err != nil {
		logger.Error("error to init database", err)
		return
	}

	queue.InitQueue(db)

}
