package common

import (
	"github.com/gin-gonic/gin"
)

func NewExceptionResponse(code int, message string) *gin.H {
	return &gin.H{
		"code":    code,
		"message": message,
	}
}

func BadRequestExceptionResponse() *gin.H {
	return NewExceptionResponse(1, "잘못된 요청입니다.")
}
