package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Module struct {
	Router *gin.RouterGroup
}

func (module *Module) AddWelcomeController() {
	module.Router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, world!",
		})
	})
}
