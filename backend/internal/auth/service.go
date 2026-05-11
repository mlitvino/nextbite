package auth

import (
	"context"

	"github.com/mlitvino/nextbite/backend/internal/models"
	"github.com/mlitvino/nextbite/backend/internal/store"
)

type Service struct {
	store store.AuthStore
}

func NewService(store store.AuthStore) *Service {
	return &Service{store: store}
}

func (s *Service) StoreCredentials(ctx context.Context, user models.User, password string) error {
	return s.store.StoreCredentials(ctx, user, password)
}

func (s *Service) Authenticate(ctx context.Context, email, password string) (models.User, error) {
	return s.store.Authenticate(ctx, email, password)
}

func (s *Service) CreateSession(ctx context.Context, userID string) (string, error) {
	return s.store.CreateSession(ctx, userID)
}

func (s *Service) GetUserBySession(ctx context.Context, token string) (models.User, error) {
	return s.store.GetUserBySession(ctx, token)
}

func (s *Service) DeleteSession(ctx context.Context, token string) error {
	return s.store.DeleteSession(ctx, token)
}
