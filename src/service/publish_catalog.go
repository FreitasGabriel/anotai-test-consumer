package service

import (
	"os"

	"github.com/FreitasGabriel/anotai-test-consumer/src/configuration/logger"
	"go.uber.org/zap"
)

func (cs *catalogServiceInterface) PublishCatalog() {
	catalog, err := cs.repo.GenerateCatalog("1")
	if err != nil {
		logger.Error("error to generate catalog", err, zap.String("journey", "publishCatalog"))
		return
	}

	err = cs.GenerateCatalogJSON(catalog)
	if err != nil {
		logger.Error("error to generate catalog JSON", err, zap.String("journey", "publishCatalog"))
	}
}

func (cs *catalogServiceInterface) GenerateCatalogJSON(catalog []byte) error {

	err := os.WriteFile("src/catalog/catalog.json", catalog, os.ModePerm)
	if err != nil {
		logger.Error("error to write JSON file", err, zap.String("journey", "generateCatalogJSON"))
		return err
	}

	return nil
}
