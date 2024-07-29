package middlewares

import (
	"errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/claim"
	"github.com/tapiaw38/cardon-tour-be/internal/platform/config"

	"github.com/gin-gonic/gin"
)

func DecodeToken(tokenString, secret string) (*domain.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &domain.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, errors.New("invalid token")
	}

	if claims, ok := token.Claims.(*domain.Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func CheckAuthMiddleware(config config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := strings.TrimSpace(c.GetHeader("Authorization"))
		claims, err := DecodeToken(tokenString, config.JWTSecret)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}
