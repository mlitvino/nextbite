package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mlitvino/nextbite/backend/internal/models"
)

type loginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

const (
	cookieName          = "nextbite_session"
	cookieMaxAgeSeconds = 3600
	loginRequestKey     = "loginRequest"
)

func (h *Handler) PostSignup(c *gin.Context) {
	h.handleSignup(c)
}

func (h *Handler) PostLogin(c *gin.Context) {
	req, ok := getLoginRequest(c)
	if !ok {
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

func getLoginRequest(c *gin.Context) (loginRequest, bool) {
	value, ok := c.Get(loginRequestKey)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "missing login request"})
		return loginRequest{}, false
	}

	req, ok := value.(loginRequest)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid login request"})
		return loginRequest{}, false
	}

	return req, true
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
