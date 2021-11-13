package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xStrato/rest-api-30min/services"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		var header string
		if header = c.GetHeader("Authorization"); header == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		token := header[len(BEARER_SCHEMA):]
		if !services.NewJWTService().Validate(token) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
