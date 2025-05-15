package middleware

import (
	"github.com/bakhytzhanjzz/go-leetcode-platform/gateway/internal/grpcclient"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware(userClient *grpcclient.UserClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid token"})
			return
		}
		token := strings.TrimPrefix(authHeader, "Bearer ")

		resp, err := userClient.ValidateToken(token)
		if err != nil || !resp.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		// store user info in context
		c.Set("user_id", resp.UserId)
		c.Next()
	}
}
