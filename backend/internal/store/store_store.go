package store

import (
	"context"
	"errors"

	"github.com/mlitvino/nextbite/backend/internal/models"
)

var ErrStoreNotFound = errors.New("store not found")

type StoreStore interface {
	List(ctx context.Context) ([]models.Store, error)
	Create(ctx context.Context, item models.Store) (models.Store, error)
	GetByID(ctx context.Context, id string) (models.Store, error)
}
