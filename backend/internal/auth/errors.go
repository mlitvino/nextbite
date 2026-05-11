package auth

import store "github.com/mlitvino/nextbite/backend/internal/repository"

var (
	ErrInvalidCredentials = store.ErrInvalidCredentials
	ErrSessionNotFound    = store.ErrSessionNotFound
	ErrUserNotFound       = store.ErrUserNotFound
	ErrUserExists         = store.ErrUserExists
)
