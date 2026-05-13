package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mlitvino/nextbite/backend/internal/models"
)

const (
	defaultRecommendationLimit = 20
	maxRecommendationLimit     = 100
)

func (h *Handler) GetMyRecommendations(c *gin.Context) {
	userValue, ok := c.Get("user")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user missing"})
		return
	}

	user, ok := userValue.(models.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user"})
		return
	}

	limit := defaultRecommendationLimit
	if raw := c.Query("limit"); raw != "" {
		parsed, err := strconv.Atoi(raw)
		if err != nil || parsed <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "limit must be a positive integer"})
			return
		}
		if parsed > maxRecommendationLimit {
			parsed = maxRecommendationLimit
		}
		limit = parsed
	}

	stores, err := h.stores.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list stores"})
		return
	}

	recommendations := h.recommender.Recommend(user, stores, limit)
	c.JSON(http.StatusOK, gin.H{"items": recommendations})
}
