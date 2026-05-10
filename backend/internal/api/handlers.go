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

type createUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (h *Handler) PostUsers(c *gin.Context) {
	var req createUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
		return
	}
	if req.Name == "" || req.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name and email are required"})
		return
	}

	created, err := h.users.Create(c.Request.Context(), req.Name, req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, created)
}
