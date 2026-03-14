package controllers

import (
	"context"
	"resource-management/internal/domain/interfaces"
)

type ResourceController struct {
	repo interfaces.ResourceRepository
}

func NewResourceController(repo interfaces.ResourceRepository) *ResourceController {
	return &ResourceController{
		repo: repo,
	}
}

func (c *ResourceController) GetResourcesCount(ctx context.Context) (int32, error) {
	return 0, nil
}
