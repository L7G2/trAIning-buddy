package middleware

import (
	"backend/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}
		token := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := utils.ParseTokenJWT(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}
		c.Set("userID", claims.UserID)
		c.Set("role", claims.Role)
		c.Next()
	}
}
