package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/mjaliz/gotracktime/internal/helpers"
	"net/http"
	"strings"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		tokenString := strings.TrimPrefix(token, "Bearer ")
		if tokenString == "" {
			helpers.FailedResponse(c, http.StatusUnauthorized, nil, "")
			c.Abort()
			return
		}
		err := helpers.ValidateToken(tokenString)
		if err != nil {
			helpers.FailedResponse(c, http.StatusUnauthorized, nil, "")
			c.Abort()
			return
		}
		c.Next()
	}
}
