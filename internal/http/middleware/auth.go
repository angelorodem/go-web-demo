package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const MockValidJWT = "MOCK_VALID_JWT"

// RequireBearer checks Authorization: Bearer <token> and compares against expected.
func RequireBearer(expected string) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		parts := strings.SplitN(auth, " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") || strings.TrimSpace(parts[1]) != expected {
			c.Header("WWW-Authenticate", `Bearer realm="api", charset="UTF-8"`)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		c.Next()
	}
}

// Convenience wrapper for your mock token.
func RequireMockToken() gin.HandlerFunc {
	return RequireBearer(MockValidJWT)
}
