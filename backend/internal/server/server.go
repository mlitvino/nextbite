package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mlitvino/nextbite/backend/internal/api"
	"github.com/mlitvino/nextbite/backend/internal/api/handlers"
	"github.com/mlitvino/nextbite/backend/internal/config"
	"github.com/mlitvino/nextbite/backend/internal/repository"
	"github.com/mlitvino/nextbite/backend/internal/repository/memory"
)

type Server struct {
	router *gin.Engine
}

func New(cfg config.Config) *Server {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	userStore := store.NewMemoryUserRepository()
	authStore := memory.NewMemoryAuthRepository()
	storeRepo := memory.NewStoreRepository()
	if err := storeRepo.LoadFromCSV(cfg.StoreSeedCSV); err != nil {
		log.Printf("store seed load failed: %v", err)
	}

	handler := handlers.NewHandler(userStore, authStore, storeRepo)
	api.RegisterRoutes(router, handler)

	return &Server{router: router}
}

func (s *Server) Run(addr string) error {
	return s.router.Run(addr)
}

func (s *Server) Router() *gin.Engine {
	return s.router
}
