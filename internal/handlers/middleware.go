package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const userCtx = "user_id"

func (h *Handler) UserIdentity(c *gin.Context) {
	header := c.GetHeader("Authorization")

	if header == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Authorization header is missing"})
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Authorization header is invalid"})
		return
	}

	userId, err := h.service.ParseToken(headerParts[1])
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Authorization header is invalid"})
		return
	}

	c.Set(userCtx, userId)
}

func (h *Handler) GetUserIdFromToken(c *gin.Context) (int, error) {
	userId, ok := c.Get("user_id")
	if !ok {
		return 0, fmt.Errorf("key not found")
	}
	id, ok := userId.(int)
	if !ok {
		return 0, fmt.Errorf("key not int")
	}
	return id, nil
}
