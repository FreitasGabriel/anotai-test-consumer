package repository

import (
	"github.com/FreitasGabriel/anotai-test-consumer/src/repository/model"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewCatalogRepository(database *mongo.Database) CatalogRepositoryInterface {
	return &catalogRepositoryInterface{
		database,
	}
}

type catalogRepositoryInterface struct {
	databaseConn *mongo.Database
}

type CatalogRepositoryInterface interface {
	FindCatalog(ownerId string) (*model.Catalog, error)
}
