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

		accessTokenClaims := entities.NewClaims(
			"access_token",
			user.Id(),
			time.Now().Add(time.Minute*30),
		)

		accessToken, err := providers.IssueAccessToken(accessTokenClaims)

		if err != nil {
			common.AbortWithException(c, FailedToIssueAccessTokenException())

			return
		}

		refreshTokenClaims := entities.NewClaims(
			"refresh_token",
			user.Id(),
			time.Now().Add(time.Hour*24*365),
		)

		refreshToken, err := providers.IssueRefreshToken(refreshTokenClaims)

		if err != nil {
			common.AbortWithException(c, FailedToIssueRefreshTokenException())

			return
		}

		user.UpdateRefreshToken(refreshToken)

		models.SaveUser(mo.Db, user)

		c.JSON(http.StatusCreated, &SignInUseCaseResponseBody{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		})
	})
}
