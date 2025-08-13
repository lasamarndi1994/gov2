package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/lasamarndi1994/gov2/internal/config"
	"github.com/lasamarndi1994/gov2/utility/response"
)

var cfg = config.LoadConfig()            //Load configuration from .env
var jwtSecret = []byte(cfg.JWTSecretKey) // 🔐 change to env var in prod

func AuthMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, response.ErrorMessage("token", "Please send Authorization token"))
		c.Abort()
		return
	}
	requestHeader := strings.Split(authHeader, " ")

	if len(requestHeader) > 2 && requestHeader[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, response.ErrorMessage("token", "Invalid token format"))
		c.Abort()
		return
	}
	tokenString := requestHeader[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, response.ErrorMessage("token", "Unauthorozied"))
		c.Abort()
		return
	}
	// Extract claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// Example: store user ID from token
		userID := uint(claims["user_id"].(float64))
		c.Set("userID", userID)
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		c.Abort()
		return
	}
	c.Next()
}
