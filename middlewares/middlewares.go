package middlewares

import (
	"net/http"

	"jwt-gin/utils/token"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValid(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"succes":  false,
				"message": "Unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}
