package common

import "github.com/gin-gonic/gin"

func AbortWithException(c *gin.Context, e Exception) {
	c.AbortWithStatusJSON(e.Status(), gin.H{
		"code":    e.Code(),
		"message": e.Message(),
	})
}
