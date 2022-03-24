package users

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/oh-jinsu/helloworld/models"
	"github.com/oh-jinsu/helloworld/modules/common"
)

type createUserRequestBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CreateUserResponseBody struct {
	Id        uint      `json:"id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
}

func (mo *Module) AddCreateUserController() {
	mo.Router.POST("", func(c *gin.Context) {
		body := &createUserRequestBody{}

		if err := c.ShouldBindJSON(body); err != nil {
			exception := common.BadRequestException()

			common.AbortWithException(c, exception)

			return
		}

		result := &models.User{}

		if err := mo.DB.Where("username = ?", body.Username).First(result).Error; err == nil {
			exception := ConflictUsernameException()

			common.AbortWithException(c, exception)

			return
		}

		mo.DB.Create(&models.User{Username: body.Username, Password: body.Password})

		mo.DB.Where("username = ?", body.Username).First(result)

		c.JSON(http.StatusCreated, &CreateUserResponseBody{
			Id:        result.ID,
			Username:  result.Username,
			CreatedAt: result.CreatedAt,
		})
	})
}
