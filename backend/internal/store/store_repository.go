package store

import (
	"context"
	"errors"

	"github.com/mlitvino/nextbite/backend/internal/models"
)

var ErrStoreNotFound = errors.New("store not found")

type StoreRepository interface {
	List(ctx context.Context) ([]models.Store, error)
	Create(ctx context.Context, item models.Store) (models.Store, error)
	GetByID(ctx context.Context, id string) (models.Store, error)
	Update(ctx context.Context, id string, item models.Store) (models.Store, error)
	Delete(ctx context.Context, id string) error
}
