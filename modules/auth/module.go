package auth

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Module struct {
	Db     *gorm.DB
	Router *gin.RouterGroup
}
