package store

import (
	"context"

	"github.com/mlitvino/nextbite/backend/internal/models"
)

type UserStore interface {
	List(ctx context.Context) ([]models.User, error)
}

type MemoryUserStore struct {
	users []models.User
}

func NewMemoryUserStore() *MemoryUserStore {
	return &MemoryUserStore{users: []models.User{}}
}

func (s *MemoryUserStore) List(ctx context.Context) ([]models.User, error) {
	_ = ctx
	return s.users, nil
}
