package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BindJSON[T any](key string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req T
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
			c.Abort()
			return
		}
		c.Set(key, req)
		c.Next()
	}
}

func RequireParam(param string, key string) gin.HandlerFunc {
	return func(c *gin.Context) {
		value := c.Param(param)
		if value == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": param + " is required"})
			c.Abort()
			return
		}
		c.Set(key, value)
		c.Next()
	}
}
