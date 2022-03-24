package users

import (
	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

type Module struct {
	DB     *gorm.DB
	Router *gin.RouterGroup
}
