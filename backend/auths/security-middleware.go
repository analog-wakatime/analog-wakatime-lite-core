package auths

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SecurityMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Cache-Control", "no-store")
		c.Header("Pragma", "no-cache")
		c.Header("Referrer-Policy", "no-referrer")
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("X-Frame-Options", "DENY")

		if hasClientSuppliedUserID(c) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Do not send user_id from client headers or query params"})
			return
		}

		c.Next()
	}
}
