package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/teixeiragthiago/api-user/internal/util"
)

type AuthenticationMiddleware interface {
	ValidateJWT() gin.HandlerFunc
}

type authMiddleware struct {
	jwtService util.JwtGeneratorService
}

func NewAuthMiddleware(jwtService util.JwtGeneratorService) AuthenticationMiddleware {
	return &authMiddleware{
		jwtService: jwtService,
	}
}

func (m *authMiddleware) ValidateJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token not informed"})
			c.Abort()
			return
		}

		tokenString := strings.TrimSpace(strings.Replace(authHeader, "Bearer", "", 1))
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			c.Abort()
			return
		}

		claims, err := m.jwtService.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		c.Set("id", claims.ID)
		c.Set("nickname", claims.Nickname)
		c.Set("email", claims.Email)
		c.Next()
	}
}
