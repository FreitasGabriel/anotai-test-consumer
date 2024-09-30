package repository

import (
	"context"

	"github.com/FreitasGabriel/anotai-test-consumer/src/configuration/logger"
	"github.com/FreitasGabriel/anotai-test-consumer/src/repository/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

var (
	ctx                context.Context
	categoryCollection = "category"
	productCollection  = "product"
)

type CategoryCatalog struct {
	ID          string              `bson:"id,omitempty"`
	Title       string              `bson:"title"`
	Description string              `bson:"description"`
	OwnerID     string              `bson:"owner_id"`
	ProductList []model.CatalogItem `bson:"products"`
}

func (cr *catalogRepositoryInterface) FindCatalog(ownerId string) (*model.CatalogPayload, error) {

	var catalog model.CatalogPayload
	var categoryItem model.CatalogCategory
	var cata []CategoryCatalog

	lookupStage := bson.D{{"$lookup", bson.D{{"from", "product"}, {"localField", "id"}, {"foreignField", "category_id"}, {"as", "products"}}}}

	cursorAgre, err := cr.databaseConn.Collection(categoryCollection).Aggregate(ctx, mongo.Pipeline{lookupStage})
	if err != nil {
		logger.Error("error to get agregation", err)
		return nil, err
	}

	defer cursorAgre.Close(ctx)

	if cursorErr := cursorAgre.All(ctx, &cata); cursorErr != nil {
		logger.Error("could not possible to cursor into slice", err, zap.String("journey", "findProductByTitle"))
		return nil, err
	}

	for _, cat := range cata {
		categoryItem.CategoryTitle = cat.Title
		categoryItem.CategoryDescription = cat.Description
		categoryItem.Itens = cat.ProductList
		catalog.Catalog = append(catalog.Catalog, categoryItem)
	}

	catalog.Owner = "1"

	return &catalog, nil
}
