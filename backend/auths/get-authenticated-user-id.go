package auths

import "github.com/gin-gonic/gin"

func GetAuthenticatedUserID(c *gin.Context) (uint, bool) {
	rawUserID, exists := c.Get(authUserIDContextKey)
	if !exists {
		return 0, false
	}

	userID, ok := rawUserID.(uint)
	return userID, ok
}
