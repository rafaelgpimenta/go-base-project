package interfaces

import "resource-management/internal/domain/entities"

// Repository interface for resource
type ResourceRepository interface {
	GetByID(id string) (*entities.Resource, error)
	Upsert(r *entities.Resource) error
	Count() (int32, error)
}
