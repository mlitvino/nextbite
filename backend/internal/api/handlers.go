package api

import (
	"github.com/mlitvino/nextbite/backend/internal/repository"
)

type Handler struct {
	users  store.UserRepository
	auth   store.AuthRepository
	stores store.StoreRepository
}

func NewHandler(users store.UserRepository, auth store.AuthRepository, stores store.StoreRepository) *Handler {
	return &Handler{users: users, auth: auth, stores: stores}
}
