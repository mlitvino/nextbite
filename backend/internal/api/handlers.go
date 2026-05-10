package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mlitvino/nextbite/backend/internal/store"
)

type Handler struct {
	users store.UserStore
}

func NewHandler(users store.UserStore) *Handler {
	return &Handler{users: users}
}

func (h *Handler) GetUsers(c *gin.Context) {
	items, err := h.users.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"items": items})
}
