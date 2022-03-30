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
			c.AbortWithStatusJSON(http.StatusBadRequest, common.BadRequestExceptionResponse())

			return
		}

		username, exception := entities.NewUsername(body.Username)

		if exception.Occured() {
			c.AbortWithStatusJSON(http.StatusBadRequest, common.NewExceptionResponse(signInExceptionCode+1, exception.Message()))

			return
		}

		user, err := models.FindUserByUsername(mo.Db, username)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, common.NewExceptionResponse(signInExceptionCode+2, "이용자를 찾지 못했습니다"))

			return
		}

		password, exception := entities.NewPassword(body.Password)

		if exception.Occured() {
			c.AbortWithStatusJSON(http.StatusBadRequest, common.NewExceptionResponse(signInExceptionCode+2, exception.Message()))

			return
		}

		if !password.Equals(user.Password()) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, common.NewExceptionResponse(signInExceptionCode+3, "비밀번호가 틀립니다"))

			return
		}

		accessTokenClaims := entities.NewClaims("access_token", user.Id(), time.Now().Add(time.Minute*30))

		accessToken, err := providers.IssueAccessToken(accessTokenClaims)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, common.NewExceptionResponse(signInExceptionCode+4, "인증 정보를 발급하지 못했습니다"))

			return
		}

		refreshTokenClaims := entities.NewClaims("refresh_token", user.Id(), time.Now().Add(time.Hour*24*365))

		refreshToken, err := providers.IssueRefreshToken(refreshTokenClaims)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, common.NewExceptionResponse(signInExceptionCode+5, "인증 정보를 발급하지 못했습니다"))

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
