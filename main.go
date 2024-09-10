package main

import (
	"github.com/FreitasGabriel/anotai-test-consumer/src/configuration/logger"
	"github.com/FreitasGabriel/anotai-test-consumer/src/service"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		logger.Error(
			"error to initialize environment variables",
			err,
		)
		return
	}

	// _, err := database.InitDatabase(context.Background())
	// if err != nil {
	// 	logger.Error("error to init database", err)
	// 	return
	// }

	service.InitQueue()
}
