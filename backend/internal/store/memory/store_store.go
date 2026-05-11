package memory

import (
	"context"
	"sync"

	"github.com/google/uuid"
	"github.com/mlitvino/nextbite/backend/internal/models"
	"github.com/mlitvino/nextbite/backend/internal/store"
)

type StoreStore struct {
	mu     sync.RWMutex
	stores []models.Store
}

func NewMemoryStoreStore() *StoreStore {
	return &StoreStore{stores: []models.Store{}}
}

func (s *StoreStore) List(ctx context.Context) ([]models.Store, error) {
	_ = ctx
	s.mu.RLock()
	defer s.mu.RUnlock()
	return append([]models.Store(nil), s.stores...), nil
}

func (s *StoreStore) Create(ctx context.Context, item models.Store) (models.Store, error) {
	_ = ctx
	if item.ID == "" {
		item.ID = uuid.NewString()
	}
	s.mu.Lock()
	s.stores = append(s.stores, item)
	s.mu.Unlock()
	return item, nil
}

func (s *StoreStore) GetByID(ctx context.Context, id string) (models.Store, error) {
	_ = ctx
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, item := range s.stores {
		if item.ID == id {
			return item, nil
		}
	}
	return models.Store{}, store.ErrStoreNotFound
}
