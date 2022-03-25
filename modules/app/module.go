package app

import (
	"github.com/gin-gonic/gin"
)

type Module struct {
	Router *gin.RouterGroup
}
