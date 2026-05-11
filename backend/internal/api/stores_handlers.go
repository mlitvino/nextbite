package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mlitvino/nextbite/backend/internal/models"
	"github.com/mlitvino/nextbite/backend/internal/store"
)

type storeRequest struct {
	Name           string     `json:"name"`
	PrimaryCuisine string     `json:"primary_cuisine"`
	Cuisines       []string   `json:"cuisines"`
	PriceTier      int        `json:"price_tier"`
	RatingAvg      float64    `json:"rating_avg"`
	RatingCount    int        `json:"rating_count"`
	Orders7D       int        `json:"orders_7d"`
	IsOpenNow      bool       `json:"is_open_now"`
	CreatedAt      *time.Time `json:"created_at"`
	Geo            models.Geo `json:"geo"`
}

func (h *Handler) GetStores(c *gin.Context) {
	items, err := h.stores.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list stores"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"items": items})
}

func (h *Handler) GetStoreByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	item, err := h.stores.GetByID(c.Request.Context(), id)
	if err != nil {
		if err == store.ErrStoreNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "store not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch store"})
		return
	}

	c.JSON(http.StatusOK, item)
}

func (h *Handler) PostStores(c *gin.Context) {
	var req storeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
		return
	}
	if req.Name == "" || req.PrimaryCuisine == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name and primary_cuisine are required"})
		return
	}

	createdAt := time.Now()
	if req.CreatedAt != nil {
		createdAt = *req.CreatedAt
	}

	created, err := h.stores.Create(c.Request.Context(), models.Store{
		Name:           req.Name,
		PrimaryCuisine: req.PrimaryCuisine,
		Cuisines:       req.Cuisines,
		PriceTier:      req.PriceTier,
		RatingAvg:      req.RatingAvg,
		RatingCount:    req.RatingCount,
		Orders7D:       req.Orders7D,
		IsOpenNow:      req.IsOpenNow,
		CreatedAt:      createdAt,
		Geo:            req.Geo,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create store"})
		return
	}

	c.JSON(http.StatusCreated, created)
}

func (h *Handler) PutStore(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	var req storeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
		return
	}
	if req.Name == "" || req.PrimaryCuisine == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name and primary_cuisine are required"})
		return
	}

	createdAt := time.Now()
	if req.CreatedAt != nil {
		createdAt = *req.CreatedAt
	} else {
		existing, err := h.stores.GetByID(c.Request.Context(), id)
		if err != nil {
			if err == store.ErrStoreNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "store not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch store"})
			return
		}
		createdAt = existing.CreatedAt
	}

	updated, err := h.stores.Update(c.Request.Context(), id, models.Store{
		Name:           req.Name,
		PrimaryCuisine: req.PrimaryCuisine,
		Cuisines:       req.Cuisines,
		PriceTier:      req.PriceTier,
		RatingAvg:      req.RatingAvg,
		RatingCount:    req.RatingCount,
		Orders7D:       req.Orders7D,
		IsOpenNow:      req.IsOpenNow,
		CreatedAt:      createdAt,
		Geo:            req.Geo,
	})
	if err != nil {
		if err == store.ErrStoreNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "store not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update store"})
		return
	}

	c.JSON(http.StatusOK, updated)
}

func (h *Handler) DeleteStore(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	if err := h.stores.Delete(c.Request.Context(), id); err != nil {
		if err == store.ErrStoreNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "store not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete store"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
