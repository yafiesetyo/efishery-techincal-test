package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	var logFormatter gin.LogFormatter = func(params gin.LogFormatterParams) string {
		var statusColor, methodColor, resetColor string
		if params.IsOutputColor() {
			statusColor = params.StatusCodeColor()
			methodColor = params.MethodColor()
			resetColor = params.ResetColor()
		}

		if params.Latency > time.Minute {
			params.Latency = params.Latency - params.Latency%time.Second
		}

		return fmt.Sprintf("[GIN] %v | %v |%s %3d %s| %13v | %15s |%s %-7s %s %#v \n%s",
			params.TimeStamp.Format("2006/01/02 - 15:04:05"),
			params.Request.Header.Get(""),
			statusColor, params.StatusCode, resetColor,
			params.Latency,
			params.ClientIP,
			methodColor, params.Method, resetColor,
			params.Path,
			params.ErrorMessage,
		)
	}
	return gin.LoggerWithFormatter(logFormatter)
}
