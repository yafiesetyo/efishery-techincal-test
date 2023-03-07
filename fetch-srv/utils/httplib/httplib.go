package httplib

import (
	"github.com/gin-gonic/gin"
)

func setErrorResponse(code int, message string) DefaultResponse {
	return DefaultResponse{
		Code:    code,
		Data:    nil,
		Message: message,
	}
}

func setSuccessResponse(code int, message string, data interface{}) DefaultResponse {
	return DefaultResponse{
		Code:    code,
		Data:    data,
		Message: message,
	}
}

func WriteResponse(c *gin.Context, code int, msg string, data interface{}) {
	if code >= 200 && code <= 299 {
		c.JSON(code, setSuccessResponse(code, msg, data))
		return
	} else {
		c.JSON(code, setErrorResponse(code, msg))
		return
	}
}
