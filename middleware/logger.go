package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func GetLogger() gin.HandlerFunc {
	logger := gin.LoggerWithFormatter(logFormat)
	return logger
}

func logFormat(param gin.LogFormatterParams) string {
	return fmt.Sprintf("%s - [%s] \"%s %s %d %s %s %s\" \n",
		param.ClientIP,
		param.TimeStamp.Format(time.RFC1123),
		param.Method,
		param.Path,
		param.StatusCode,
		param.Latency,
		param.Request.UserAgent(),
		param.ErrorMessage,
	)
}
