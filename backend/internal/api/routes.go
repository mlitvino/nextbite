package api

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine, h *Handler) {
	api := r.Group("/api")
	api.GET("/health", health)
	api.GET("/users", h.GetUsers)
	api.POST("/users", h.PostUsers)
	api.POST("/login", h.PostLogin)
	api.POST("/logout", h.PostLogout)
	api.GET("/me", h.RequireAuth(), h.GetMe)
}
