package api

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine, h *Handler) {
	api := r.Group("/api")
	registerHealthRoute(api)
	registerUsersRoutes(api, h)
	registerStoresRoutes(api, h)
	registerAuthRoutes(api, h)
}

func registerHealthRoute(r gin.IRoutes) {
	r.GET("/health", health)
}

func registerUsersRoutes(r gin.IRoutes, h *Handler) {
	r.GET("/users", h.GetUsers)
	r.POST("/users", h.PostUsers)
}

func registerStoresRoutes(r gin.IRoutes, h *Handler) {
	r.GET("/stores", h.GetStores)
	r.POST("/stores", h.PostStores)
	r.GET("/stores/:id", h.GetStoreByID)
	r.PUT("/stores/:id", h.PutStore)
	r.DELETE("/stores/:id", h.DeleteStore)
}

func registerAuthRoutes(r gin.IRoutes, h *Handler) {
	r.POST("/auth/signup", h.PostSignup)
	r.POST("/auth/login", h.PostLogin)
	r.POST("/auth/logout", h.PostLogout)
	r.GET("/me", h.RequireAuth(), h.GetMe)
}
