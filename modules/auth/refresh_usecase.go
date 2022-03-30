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
			c.AbortWithStatusJSON(http.StatusBadRequest, common.BadRequestExceptionResponse())

			return
		}

		claims, err := providers.VerifyRefreshToken(body.RefreshToken)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, common.NewExceptionResponse(refreshExceptionCode+1, "유효하지 않은 인증 정보입니다"))

			return
		}

		user, err := models.FindUserById(mo.Db, claims.UserId())

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, common.NewExceptionResponse(refreshExceptionCode+2, "이용자를 찾지 못했습니다"))

			return
		}

		if user.RefreshToken() != body.RefreshToken {
			c.AbortWithStatusJSON(http.StatusUnauthorized, common.NewExceptionResponse(refreshExceptionCode+3, "이미 폐기된 인증 정보입니다"))

			return
		}

		accessTokenClaims := entities.NewClaims(
			"access_token",
			user.Id(),
			time.Now().Add(time.Minute*30),
		)

		accessToken, err := providers.IssueAccessToken(accessTokenClaims)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, common.NewExceptionResponse(refreshExceptionCode+4, "인증 정보를 발급하지 못했습니다"))

			return
		}

		c.JSON(http.StatusCreated, &RefreshUseCaseResponseBody{
			AccessToken: accessToken,
		})
	})
}
