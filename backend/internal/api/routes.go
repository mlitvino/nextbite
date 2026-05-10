package api

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine, h *Handler) {
	api := r.Group("/api")
	api.GET("/health", health)
	api.GET("/users", h.GetUsers)
	api.POST("/users", h.PostUsers)
}
