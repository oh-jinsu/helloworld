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

type refreshUseCaseRequestBody struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type RefreshUseCaseResponseBody struct {
	AccessToken string `json:"access_token"`
}

func (mo *Module) AddRefreshUseCase() {
	mo.Router.POST("refresh", func(c *gin.Context) {
		body := &refreshUseCaseRequestBody{}

		if err := c.ShouldBindJSON(body); err != nil {
			common.AbortWithException(c, common.BadRequestException())

			return
		}

		claims, err := providers.VerifyRefreshToken(body.RefreshToken)

		if err != nil {
			common.AbortWithException(c, InvalidRefreshTokenException())

			return
		}

		user, err := models.FindUser(mo.Db, claims.UserId())

		if err != nil {
			common.AbortWithException(c, UserNotFoundException())

			return
		}

		if user.RefreshToken() != body.RefreshToken {
			common.AbortWithException(c, DiscardedRefreshTokenException())

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

		c.JSON(http.StatusCreated, &RefreshUseCaseResponseBody{
			AccessToken: accessToken,
		})
	})
}
