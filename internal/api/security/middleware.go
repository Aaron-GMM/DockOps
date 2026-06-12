package security

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := ParseToken(c.Request.Header, secret)
		if err != nil {
			log.WarningF("Acesso negado pela camada de segurança: %v", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"Error":  "Unauthorized",
				"code":   http.StatusUnauthorized,
				"detail": err.Error(),
			})
			return
		}
		if sub, ok := claims["sub"].(string); ok {
			c.Set("userID", sub)
		}
		if role, ok := claims["role"].(string); ok {
			c.Set("role", role)
		}

		c.Next()
	}

}
