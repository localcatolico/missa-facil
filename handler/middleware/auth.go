package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rafaelbmateus/slides-gospel/handler/web"
)

// AuthRequired is a simple middleware to check the session
func AuthRequired(c *gin.Context) {
	user := web.UserToken(c)
	if user == nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	c.Next()
}
