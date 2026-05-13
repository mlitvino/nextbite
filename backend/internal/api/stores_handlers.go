package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mlitvino/nextbite/backend/internal/models"
	store "github.com/mlitvino/nextbite/backend/internal/repository"
)

type storeRequest struct {
	Name           string     `json:"name" binding:"required"`
	PrimaryCuisine string     `json:"primary_cuisine" binding:"required"`
	Cuisines       []string   `json:"cuisines"`
	PriceTier      int        `json:"price_tier"`
	RatingAvg      float64    `json:"rating_avg"`
	RatingCount    int        `json:"rating_count"`
	Orders7D       int        `json:"orders_7d"`
	IsOpenNow      bool       `json:"is_open_now"`
	CreatedAt      *time.Time `json:"created_at"`
	Geo            models.Geo `json:"geo"`
}

const (
	storeRequestKey = "storeRequest"
	storeIDKey      = "storeID"
)

func (h *Handler) GetStores(c *gin.Context) {
	items, err := h.stores.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list stores"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"items": items})
}

func (h *Handler) GetStoreByID(c *gin.Context) {
	id, ok := getStoreID(c)
	if !ok {
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
	req, ok := getStoreRequest(c)
	if !ok {
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
	id, ok := getStoreID(c)
	if !ok {
		return
	}

	req, ok := getStoreRequest(c)
	if !ok {
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
	id, ok := getStoreID(c)
	if !ok {
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

func getStoreRequest(c *gin.Context) (storeRequest, bool) {
	value, ok := c.Get(storeRequestKey)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "missing store request"})
		return storeRequest{}, false
	}
	req, ok := value.(storeRequest)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid store request"})
		return storeRequest{}, false
	}
	return req, true
}

func getStoreID(c *gin.Context) (string, bool) {
	value, ok := c.Get(storeIDKey)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "missing store id"})
		return "", false
	}
	id, ok := value.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid store id"})
		return "", false
	}
	return id, true
}
