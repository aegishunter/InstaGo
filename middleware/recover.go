package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRecoverHandler() gin.HandlerFunc {
	return gin.CustomRecovery(recoveryHandler)
}

func recoveryHandler(c *gin.Context, recovered interface{}) {
	if err, ok := recovered.(string); ok {
		c.String(http.StatusInternalServerError, err)
	}
	c.AbortWithStatus(http.StatusInternalServerError)
}
