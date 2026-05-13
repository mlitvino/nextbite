package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mlitvino/nextbite/backend/internal/models"
)

type createUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

const createUserRequestKey = "createUserRequest"

func (h *Handler) GetUsers(c *gin.Context) {
	items, err := h.users.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"items": items})
}

func (h *Handler) PostUsers(c *gin.Context) {
	h.handleSignup(c)
}

func (h *Handler) handleSignup(c *gin.Context) {
	req, ok := getCreateUserRequest(c)
	if !ok {
		return
	}

	created, err := h.users.Create(c.Request.Context(), models.User{
		Name:     req.Name,
		Username: req.Username,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	if err := h.auth.StoreCredentials(c.Request.Context(), created, req.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, created)
}

func getCreateUserRequest(c *gin.Context) (createUserRequest, bool) {
	value, ok := c.Get(createUserRequestKey)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "missing create user request"})
		return createUserRequest{}, false
	}

	req, ok := value.(createUserRequest)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid create user request"})
		return createUserRequest{}, false
	}

	return req, true
}
