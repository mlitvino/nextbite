package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mlitvino/nextbite/backend/internal/models"
	"github.com/mlitvino/nextbite/backend/internal/store"
)

type Handler struct {
	users store.UserRepository
	auth  store.AuthRepository
}

func NewHandler(users store.UserRepository, auth store.AuthRepository) *Handler {
	return &Handler{users: users, auth: auth}
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
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *Handler) PostUsers(c *gin.Context) {
	h.handleSignup(c)
}

func (h *Handler) PostSignup(c *gin.Context) {
	h.handleSignup(c)
}

func (h *Handler) handleSignup(c *gin.Context) {
	var req createUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
		return
	}
	if req.Name == "" || req.Username == "" || req.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name, username, and password are required"})
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

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

const (
	cookieName          = "nextbite_session"
	cookieMaxAgeSeconds = 3600
)

func (h *Handler) PostLogin(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
		return
	}
	if req.Username == "" || req.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username and password are required"})
		return
	}

	user, err := h.auth.Authenticate(c.Request.Context(), req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	token, err := h.auth.CreateSession(c.Request.Context(), user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create session"})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie(cookieName, token, cookieMaxAgeSeconds, "/", "", false, true)
	c.JSON(http.StatusOK, user)
}

func (h *Handler) PostLogout(c *gin.Context) {
	if token, err := c.Cookie(cookieName); err == nil && token != "" {
		_ = h.auth.DeleteSession(c.Request.Context(), token)
	}

	c.SetCookie(cookieName, "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *Handler) GetMe(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user missing"})
		return
	}

	current, ok := user.(models.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user"})
		return
	}

	c.JSON(http.StatusOK, current)
}

func (h *Handler) RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie(cookieName)
		if err != nil || token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		user, err := h.auth.GetUserBySession(c.Request.Context(), token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
