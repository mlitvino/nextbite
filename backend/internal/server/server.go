package server

import (
	"github.com/gin-gonic/gin"
	"github.com/mlitvino/nextbite/backend/internal/api"
	"github.com/mlitvino/nextbite/backend/internal/store"
	"github.com/mlitvino/nextbite/backend/internal/store/memory"
)

type Server struct {
	router *gin.Engine
}

func New() *Server {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	userStore := store.NewMemoryUserStore()
	authStore := memory.NewMemoryAuthStore()

	handler := api.NewHandler(userStore, authStore)
	api.RegisterRoutes(router, handler)

	return &Server{router: router}
}

func (s *Server) Run(addr string) error {
	return s.router.Run(addr)
}

func (s *Server) Router() *gin.Engine {
	return s.router
}
