package repository

import "context"

var (
	ctx      context.Context
	viewName = "catalog"
)

func (cr *catalogRepositoryInterface) FindCatalog() error {

	return nil
}

func (cr *catalogRepositoryInterface) CreateVisualization() error {
	cr.databaseConn.CreateView(ctx, viewName, "")
}
