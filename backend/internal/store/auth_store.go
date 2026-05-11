package store

import (
	"context"
	"errors"

	"github.com/mlitvino/nextbite/backend/internal/models"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrSessionNotFound    = errors.New("session not found")
	ErrUserNotFound       = errors.New("user not found")
	ErrUserExists         = errors.New("user already exists")
)

type AuthStore interface {
	Authenticate(ctx context.Context, username, password string) (models.User, error)
	CreateSession(ctx context.Context, userID string) (string, error)
	GetUserBySession(ctx context.Context, token string) (models.User, error)
	DeleteSession(ctx context.Context, token string) error
	StoreCredentials(ctx context.Context, user models.User, password string) error
}
