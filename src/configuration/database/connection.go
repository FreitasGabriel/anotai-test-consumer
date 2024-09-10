package database

import (
	"context"
	"os"

	"github.com/FreitasGabriel/anotai-test-consumer/src/configuration/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

const (
	DATABASE_URL  = "DATABASE_URL"
	DATABASE_NAME = "DATABASE_NAME"
)

func InitDatabase(ctx context.Context) (*mongo.Database, error) {
	databaseURL := os.Getenv(DATABASE_URL)
	databaseName := os.Getenv(DATABASE_NAME)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(databaseURL))
	if err != nil {
		logger.Error("error to get client", err, zap.String("journey", "getDatabaseClient"))
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		logger.Error("error to ping database", err, zap.String("journey", "pingDatabase"))
		return nil, err
	}

	return client.Database(databaseName), nil
}
