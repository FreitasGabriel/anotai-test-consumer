package service

import "github.com/FreitasGabriel/anotai-test-consumer/src/repository"

func NewCatalogService(repo repository.CatalogRepositoryInterface) CatalogServiceInterface {
	return &catalogServiceInterface{
		repo,
	}
}

type catalogServiceInterface struct {
	repo repository.CatalogRepositoryInterface
}

type CatalogServiceInterface interface {
	PublishCatalog()
}
