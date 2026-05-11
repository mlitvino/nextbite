package memory

import (
	"context"
	"sync"

	"github.com/google/uuid"
	"github.com/mlitvino/nextbite/backend/internal/models"
	"github.com/mlitvino/nextbite/backend/internal/store"
)

type StoreRepository struct {
	mu     sync.RWMutex
	stores []models.Store
}

func NewStoreRepository() *StoreRepository {
	return &StoreRepository{stores: []models.Store{}}
}

func (s *StoreRepository) List(ctx context.Context) ([]models.Store, error) {
	_ = ctx
	s.mu.RLock()
	defer s.mu.RUnlock()
	return append([]models.Store(nil), s.stores...), nil
}

func (s *StoreRepository) Create(ctx context.Context, item models.Store) (models.Store, error) {
	_ = ctx
	if item.ID == "" {
		item.ID = uuid.NewString()
	}
	s.mu.Lock()
	s.stores = append(s.stores, item)
	s.mu.Unlock()
	return item, nil
}

func (s *StoreRepository) GetByID(ctx context.Context, id string) (models.Store, error) {
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

func (s *StoreRepository) Update(ctx context.Context, id string, item models.Store) (models.Store, error) {
	_ = ctx
	s.mu.Lock()
	defer s.mu.Unlock()
	for i := range s.stores {
		if s.stores[i].ID == id {
			item.ID = id
			s.stores[i] = item
			return item, nil
		}
	}
	return models.Store{}, store.ErrStoreNotFound
}

func (s *StoreRepository) Delete(ctx context.Context, id string) error {
	_ = ctx
	s.mu.Lock()
	defer s.mu.Unlock()
	for i := range s.stores {
		if s.stores[i].ID == id {
			s.stores = append(s.stores[:i], s.stores[i+1:]...)
			return nil
		}
	}
	return store.ErrStoreNotFound
}
