package handlers

import (
	"github.com/mlitvino/nextbite/backend/internal/recommendation"
	store "github.com/mlitvino/nextbite/backend/internal/repository"
	"github.com/mlitvino/nextbite/backend/internal/scoring"
)

type Handler struct {
	users       store.UserRepository
	auth        store.AuthRepository
	stores      store.StoreRepository
	recommender *recommendation.Engine
}

func NewHandler(users store.UserRepository, auth store.AuthRepository, stores store.StoreRepository) *Handler {
	scorer := scoring.NewDefaultScorer()
	engine := recommendation.NewEngine(scorer)
	return &Handler{users: users, auth: auth, stores: stores, recommender: engine}
}
