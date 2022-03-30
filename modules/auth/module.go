package auth

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Module struct {
	Db     *gorm.DB
	Router *gin.RouterGroup
}

const (
	signUpExceptionCode = 100 + iota*100
	signInExceptionCode
	refreshExceptionCode
)
