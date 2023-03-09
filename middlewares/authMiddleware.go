package middlewares

import (
	"net/http"
	"strings"

	"github.com/ariwanss/task-5-vix-btpns-ariwan-sri-setya/helpers"
	"github.com/gin-gonic/gin"
)

func Protect(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")

	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized no auth header"})
		return
	}

	if !strings.HasPrefix(authHeader, "Bearer") {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized no bearer token"})
		return
	}

	tokenStr := strings.Split(authHeader, " ")[1]
	payload, err := helpers.VerifyToken(tokenStr)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.Set("User", payload)
	c.Next()
}
