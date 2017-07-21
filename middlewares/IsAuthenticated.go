package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/guszak/test/conn"
	"gitlab.com/guszak/test/models"
)

// IsAuthenticated control access auth users
func IsAuthenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")

		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "API token required",
			})
			return
		}

		db := conn.InitDb()
		defer db.Close()
		var company models.Company
		if err := db.Where("token = ?", token).First(&company).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid API token",
			})
			return
		}
		c.Set("company", &company)

		c.Next()
	}
}
