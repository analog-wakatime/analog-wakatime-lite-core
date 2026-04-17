package profiles

import (
	"analog-wakatime-lite-core/db"
	"analog-wakatime-lite-core/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMyInfo(c *gin.Context) {
	userID, exists := c.Get("auth_user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var user models.User
	if err := db.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user info"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":       user.ID,
		"email":    user.Email,
		"username": user.Username,
	})
}
