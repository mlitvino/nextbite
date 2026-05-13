package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mlitvino/nextbite/backend/internal/api/handlers"
)

func RegisterRoutes(r *gin.Engine, h *handlers.Handler) {
	api := r.Group("/api")
	registerHealthRoute(api)
	registerUsersRoutes(api, h)
	registerStoresRoutes(api, h)
	registerAuthRoutes(api, h)
}

func registerHealthRoute(r gin.IRoutes) {
	r.GET("/health", health)
}

func registerUsersRoutes(r gin.IRoutes, h *handlers.Handler) {
	r.GET("/users", h.GetUsers)
	r.POST("/users", BindJSON[handlers.CreateUserRequest](handlers.CreateUserRequestKey), h.PostUsers)
}

func registerStoresRoutes(r gin.IRoutes, h *handlers.Handler) {
	r.GET("/stores", h.GetStores)
	r.POST("/stores", BindJSON[handlers.StoreRequest](handlers.StoreRequestKey), h.PostStores)
	r.GET("/stores/:id", RequireParam("id", handlers.StoreIDKey), h.GetStoreByID)
	r.PUT("/stores/:id", RequireParam("id", handlers.StoreIDKey), BindJSON[handlers.StoreRequest](handlers.StoreRequestKey), h.PutStore)
	r.DELETE("/stores/:id", RequireParam("id", handlers.StoreIDKey), h.DeleteStore)
}

func registerAuthRoutes(r gin.IRoutes, h *handlers.Handler) {
	r.POST("/auth/signup", BindJSON[handlers.CreateUserRequest](handlers.CreateUserRequestKey), h.PostSignup)
	r.POST("/auth/login", BindJSON[handlers.LoginRequest](handlers.LoginRequestKey), h.PostLogin)
	r.POST("/auth/logout", h.PostLogout)
	r.GET("/me", h.RequireAuth(), h.GetMe)
}
