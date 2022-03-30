package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/oh-jinsu/helloworld/entities"
	"github.com/oh-jinsu/helloworld/models"
	"github.com/oh-jinsu/helloworld/modules/common"
)

type signUpRequestBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignUpResponseBody struct {
	Id        uint      `json:"id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
}

func (mo *Module) AddSignUpUseCase() {
	mo.Router.POST("", func(c *gin.Context) {
		body := &signUpRequestBody{}

		if err := c.ShouldBindJSON(body); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, common.BadRequestExceptionResponse())

			return
		}

		username, exception := entities.NewUsername(body.Username)

		if exception.Occured() {
			c.AbortWithStatusJSON(http.StatusBadRequest, common.NewExceptionResponse(signUpExceptionCode+1, exception.Message()))

			return
		}

		password, exception := entities.NewPassword(body.Password)

		if exception.Occured() {
			c.AbortWithStatusJSON(http.StatusBadRequest, common.NewExceptionResponse(signUpExceptionCode+2, exception.Message()))

			return
		}

		if models.ExistsByUsername(mo.Db, username) {
			c.AbortWithStatusJSON(http.StatusConflict, common.NewExceptionResponse(signUpExceptionCode+3, "인증 정보를 발급하지 못했습니다"))

			return
		}

		userId := models.NextUserId(mo.Db)

		user, exception := entities.NewUser(userId, username, password)

		if exception.Occured() {
			c.AbortWithStatusJSON(http.StatusBadRequest, common.NewExceptionResponse(signUpExceptionCode+4, exception.Message()))
		}

		models.SaveUser(mo.Db, user)

		c.JSON(http.StatusCreated, &SignUpResponseBody{
			Id:        user.Id(),
			Username:  user.Username().ToString(),
			CreatedAt: user.CreatedAt(),
		})
	})
}
