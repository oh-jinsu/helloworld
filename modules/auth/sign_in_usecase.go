package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/oh-jinsu/helloworld/entities"
	"github.com/oh-jinsu/helloworld/models"
	"github.com/oh-jinsu/helloworld/modules/common"
	"github.com/oh-jinsu/helloworld/providers"
)

type signInUseCaseRequestBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignInUseCaseResponseBody struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (mo *Module) AddSignInUseCase() {
	mo.Router.POST("issue", func(c *gin.Context) {
		body := &signInUseCaseRequestBody{}

		if err := c.ShouldBindJSON(body); err != nil {
			common.AbortWithException(c, common.BadRequestException())

			return
		}

		username := entities.NewUsername(body.Username)

		user, err := models.FindUserByUsername(mo.Db, username)

		if err != nil {
			common.AbortWithException(c, UserNotFoundException())

			return
		}

		password := entities.NewPassword(body.Password)

		if !password.Equals(user.Password()) {
			common.AbortWithException(c, PasswordNotMatchedException())

			return
		}

		accessToken, err := providers.IssueAccessToken(user.Id(), time.Minute*30)

		if err != nil {
			common.AbortWithException(c, FailedToIssueAccessTokenException())

			return
		}

		refreshToken, err := providers.IssueRefreshToken(user.Id(), time.Hour*24*365)

		if err != nil {
			common.AbortWithException(c, FailedToIssueRefreshTokenException())

			return
		}

		c.JSON(http.StatusCreated, &SignInUseCaseResponseBody{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		})
	})
}
