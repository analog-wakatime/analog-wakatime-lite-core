package auths

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func hasClientSuppliedUserID(c *gin.Context) bool {
	for _, header := range []string{"User-ID", "X-User-ID", "X-User-Id"} {
		if strings.TrimSpace(c.GetHeader(header)) != "" {
			return true
		}
	}

	return strings.TrimSpace(c.Query("user_id")) != ""
}
