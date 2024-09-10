package repository

import "go.mongodb.org/mongo-driver/mongo"

func NewCatalogRepository(database *mongo.Database) CatalogRepositoryInterface {
	return &catalogRepositoryInterface{
		database,
	}
}

type catalogRepositoryInterface struct {
	databaseConn *mongo.Database
}

type CatalogRepositoryInterface interface {
	FindCatalog() error
	CreateVisualization() error
}
