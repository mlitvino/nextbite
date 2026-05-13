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
	r.POST("/users", BindJSON[createUserRequest](createUserRequestKey), h.PostUsers)
}

func registerStoresRoutes(r gin.IRoutes, h *Handler) {
	r.GET("/stores", h.GetStores)
	r.POST("/stores", BindJSON[storeRequest](storeRequestKey), h.PostStores)
	r.GET("/stores/:id", RequireParam("id", storeIDKey), h.GetStoreByID)
	r.PUT("/stores/:id", RequireParam("id", storeIDKey), BindJSON[storeRequest](storeRequestKey), h.PutStore)
	r.DELETE("/stores/:id", RequireParam("id", storeIDKey), h.DeleteStore)
}

func registerAuthRoutes(r gin.IRoutes, h *Handler) {
	r.POST("/auth/signup", BindJSON[createUserRequest](createUserRequestKey), h.PostSignup)
	r.POST("/auth/login", BindJSON[loginRequest](loginRequestKey), h.PostLogin)
	r.POST("/auth/logout", h.PostLogout)
	r.GET("/me", h.RequireAuth(), h.GetMe)
}
