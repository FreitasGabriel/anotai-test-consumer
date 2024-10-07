package repository

import (
	"context"
	"encoding/json"

	"github.com/FreitasGabriel/anotai-test-consumer/src/configuration/logger"
	"github.com/FreitasGabriel/anotai-test-consumer/src/repository/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

var (
	ctx                context.Context
	categoryCollection = "category"
)

func (cr *catalogRepositoryInterface) GenerateCatalog(ownerId string) ([]byte, error) {

	var catalog model.Catalog
	var categoryItem model.Category
	var catalogAggregation []model.CatalogAggregation

	lookupStage := bson.D{{"$lookup", bson.D{{"from", "product"}, {"localField", "id"}, {"foreignField", "category_id"}, {"as", "products"}}}}

	cursor, err := cr.databaseConn.Collection(categoryCollection).Aggregate(ctx, mongo.Pipeline{lookupStage})
	if err != nil {
		logger.Error("error to get agregation", err, zap.String("journey", "generateCatalog"))
		return nil, err
	}

	defer cursor.Close(ctx)

	if cursorErr := cursor.All(ctx, &catalogAggregation); cursorErr != nil {
		logger.Error("could not possible to cursor into slice", err, zap.String("journey", "generateCatalog"))
		return nil, err
	}

	catalog.Owner = "1"
	for _, catalogItem := range catalogAggregation {
		categoryItem.CategoryTitle = catalogItem.Title
		categoryItem.CategoryDescription = catalogItem.Description
		categoryItem.Itens = catalogItem.ProductList
		catalog.Catalog = append(catalog.Catalog, categoryItem)
	}

	res, err := json.Marshal(catalog)
	if err != nil {
		logger.Error("error to marshal catalog", err, zap.String("journey", "generateCatalog"))
	}

	return res, nil
}
