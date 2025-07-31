package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/loid-lab/restaurant-menu/models"
)

func CheckAdmin(c *gin.Context) {
	user, exists := c.Get("currentUser")
	if !exists {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized"})
		return
	}
	role := user.(models.User).Role
	if role != "admin" && role != "staff" {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"error": "forbidden"})
	}
	c.Next()
}
