package middlewares

import (
	"net/http"

	"github.com/XanderMoroz/BookStore/utils"
	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := utils.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Header("new-token", "123")
		c.Next()
	}
}
