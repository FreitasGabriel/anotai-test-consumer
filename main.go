package main

import (
	"github.com/FreitasGabriel/anotai-test-consumer/src/configuration/logger"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	err := router.Run(":8081")
	if err != nil {
		logger.Error("error to init server", err)
		panic(err)
	}

	logger.Info("server runing in the port: 8081")

}
