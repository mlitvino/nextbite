package store

import (
	"context"

	"github.com/google/uuid"
	"github.com/mlitvino/nextbite/backend/internal/models"
)

type UserStore interface {
	List(ctx context.Context) ([]models.User, error)
	Create(ctx context.Context, user models.User) (models.User, error)
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

func (s *MemoryUserStore) Create(ctx context.Context, user models.User) (models.User, error) {
	_ = ctx
	if user.ID == "" {
		user.ID = uuid.NewString()
	}
	s.users = append(s.users, user)
	return user, nil
}
