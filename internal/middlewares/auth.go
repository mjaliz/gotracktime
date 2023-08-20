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
		authorizationHeader := c.GetHeader(constants.AuthorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			utils.FailedResponse(c, http.StatusUnauthorized, nil, "")
			c.Abort()
			return
		}
		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			utils.FailedResponse(c, http.StatusUnauthorized, nil, "")
			c.Abort()
			return
		}
		authorizationType := strings.ToLower(fields[0])
		if authorizationType != constants.AuthorizationTypeBearer {
			utils.FailedResponse(c, http.StatusUnauthorized, nil, "")
			c.Abort()
			return
		}
		accessToken := fields[1]
		payload, err := utils.ValidateToken(accessToken)
		if err != nil {
			utils.FailedResponse(c, http.StatusUnauthorized, nil, "")
			c.Abort()
			return
		}
		c.Set(constants.AuthorizationPayload, payload)
		c.Next()
	}
}
