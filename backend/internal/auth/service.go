package auth

import (
	"context"

	"github.com/mlitvino/nextbite/backend/internal/models"
	"github.com/mlitvino/nextbite/backend/internal/repository"
)

type Service struct {
	repo store.AuthRepository
}

func NewService(repo store.AuthRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) StoreCredentials(ctx context.Context, user models.User, password string) error {
	return s.repo.StoreCredentials(ctx, user, password)
}

func (s *Service) Authenticate(ctx context.Context, username, password string) (models.User, error) {
	return s.repo.Authenticate(ctx, username, password)
}

func (s *Service) CreateSession(ctx context.Context, userID string) (string, error) {
	return s.repo.CreateSession(ctx, userID)
}

func (s *Service) GetUserBySession(ctx context.Context, token string) (models.User, error) {
	return s.repo.GetUserBySession(ctx, token)
}

func (s *Service) DeleteSession(ctx context.Context, token string) error {
	return s.repo.DeleteSession(ctx, token)
}
