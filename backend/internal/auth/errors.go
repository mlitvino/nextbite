package auth

import "github.com/mlitvino/nextbite/backend/internal/store"

var (
	ErrInvalidCredentials = store.ErrInvalidCredentials
	ErrSessionNotFound    = store.ErrSessionNotFound
	ErrUserNotFound       = store.ErrUserNotFound
	ErrUserExists         = store.ErrUserExists
)
