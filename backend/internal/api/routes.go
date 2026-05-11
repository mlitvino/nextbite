package api

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine, h *Handler) {
	api := r.Group("/api")
	api.GET("/health", health)
	api.GET("/users", h.GetUsers)
	api.POST("/users", h.PostUsers)
	api.GET("/stores", h.GetStores)
	api.POST("/stores", h.PostStores)
	api.GET("/stores/:id", h.GetStoreByID)
	api.PUT("/stores/:id", h.PutStore)
	api.DELETE("/stores/:id", h.DeleteStore)
	api.POST("/auth/signup", h.PostSignup)
	api.POST("/auth/login", h.PostLogin)
	api.POST("/auth/logout", h.PostLogout)
	api.GET("/me", h.RequireAuth(), h.GetMe)
}
