package store

import (
	"context"

	"github.com/google/uuid"
	"github.com/mlitvino/nextbite/backend/internal/models"
)

type UserRepository interface {
	List(ctx context.Context) ([]models.User, error)
	Create(ctx context.Context, user models.User) (models.User, error)
}

type MemoryUserRepository struct {
	users []models.User
}

func NewMemoryUserRepository() *MemoryUserRepository {
	return &MemoryUserRepository{users: []models.User{}}
}

func (s *MemoryUserRepository) List(ctx context.Context) ([]models.User, error) {
	_ = ctx
	return s.users, nil
}

func (s *MemoryUserRepository) Create(ctx context.Context, user models.User) (models.User, error) {
	_ = ctx
	if user.ID == "" {
		user.ID = uuid.NewString()
	}
	s.users = append(s.users, user)
	return user, nil
}
