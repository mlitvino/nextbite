package server

import (
	"github.com/gin-gonic/gin"
	"github.com/mlitvino/nextbite/backend/internal/api"
	"github.com/mlitvino/nextbite/backend/internal/repository"
	"github.com/mlitvino/nextbite/backend/internal/repository/memory"
)

type Server struct {
	router *gin.Engine
}

func New() *Server {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	userStore := store.NewMemoryUserRepository()
	authStore := memory.NewMemoryAuthRepository()
	storeRepo := memory.NewStoreRepository()

	handler := api.NewHandler(userStore, authStore, storeRepo)
	api.RegisterRoutes(router, handler)

	return &Server{router: router}
}

func (s *Server) Run(addr string) error {
	return s.router.Run(addr)
}

func (s *Server) Router() *gin.Engine {
	return s.router
}
