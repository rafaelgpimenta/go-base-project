package application

import (
	"context"
	"resource-management/internal/domain/interfaces"
)

type Controller struct {
	db interfaces.Database
}

func NewController(db interfaces.Database) *Controller {
	return &Controller{
		db: db,
	}
}

func (c *Controller) GetResourcesCount(ctx context.Context) (int32, error) {
	return 0, nil
}
