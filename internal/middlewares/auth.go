package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/mjaliz/gotracktime/internal/constants"
	"github.com/mjaliz/gotracktime/internal/utils"
	"net/http"
	"strings"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		tokenString := strings.TrimPrefix(token, "Bearer ")
		if tokenString == "" {
			utils.FailedResponse(c, http.StatusUnauthorized, nil, "")
			c.Abort()
			return
		}
		claim, err := utils.ValidateToken(tokenString)
		if err != nil {
			utils.FailedResponse(c, http.StatusUnauthorized, nil, "")
			c.Abort()
			return
		}
		c.Set(constants.UserClaims, claim)
		c.Next()
	}
}
