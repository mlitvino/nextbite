package memory

import (
	"context"
	"sync"

	"github.com/google/uuid"
	"github.com/mlitvino/nextbite/backend/internal/models"
	"github.com/mlitvino/nextbite/backend/internal/store"
	"golang.org/x/crypto/bcrypt"
)

type authUser struct {
	user         models.User
	passwordHash []byte
}

type AuthStore struct {
	mu         sync.RWMutex
	byUsername map[string]authUser
	byID       map[string]models.User
	sessions   map[string]string
}

func NewMemoryAuthStore() *AuthStore {
	return &AuthStore{
		byUsername: make(map[string]authUser),
		byID:       make(map[string]models.User),
		sessions:   make(map[string]string),
	}
}

func (s *AuthStore) StoreCredentials(ctx context.Context, user models.User, password string) error {
	_ = ctx
	if user.ID == "" || user.Username == "" || password == "" {
		return store.ErrInvalidCredentials
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.byUsername[user.Username]; exists {
		return store.ErrUserExists
	}
	s.byUsername[user.Username] = authUser{user: user, passwordHash: hash}
	s.byID[user.ID] = user

	return nil
}

func (s *AuthStore) Authenticate(ctx context.Context, username, password string) (models.User, error) {
	_ = ctx
	s.mu.RLock()
	entry, ok := s.byUsername[username]
	s.mu.RUnlock()
	if !ok {
		return models.User{}, store.ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword(entry.passwordHash, []byte(password)); err != nil {
		return models.User{}, store.ErrInvalidCredentials
	}

	return entry.user, nil
}

func (s *AuthStore) CreateSession(ctx context.Context, userID string) (string, error) {
	_ = ctx
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.byID[userID]; !ok {
		return "", store.ErrUserNotFound
	}

	token := uuid.NewString()
	s.sessions[token] = userID
	return token, nil
}

func (s *AuthStore) GetUserBySession(ctx context.Context, token string) (models.User, error) {
	_ = ctx
	s.mu.RLock()
	userID, ok := s.sessions[token]
	if !ok {
		s.mu.RUnlock()
		return models.User{}, store.ErrSessionNotFound
	}
	user, ok := s.byID[userID]
	s.mu.RUnlock()
	if !ok {
		return models.User{}, store.ErrUserNotFound
	}

	return user, nil
}

func (s *AuthStore) DeleteSession(ctx context.Context, token string) error {
	_ = ctx
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.sessions, token)
	return nil
}
