package users

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/oh-jinsu/helloworld/entities"
	"github.com/oh-jinsu/helloworld/modules/common"
	"github.com/oh-jinsu/helloworld/repositories"
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

func (mo *Module) AddCreateUserUseCase() {
	mo.Router.POST("", func(c *gin.Context) {
		body := &createUserRequestBody{}

		if err := c.ShouldBindJSON(body); err != nil {
			common.AbortWithException(c, common.BadRequestException())

			return
		}

		username := entities.NewUsername(body.Username)

		if username.HasKoreanConsonants() {
			common.AbortWithException(c, KoreanConsonantsException())

			return
		}

		if username.HasSpecialCharacters() {
			common.AbortWithException(c, SpecialCharacterException())

			return
		}

		if username.HasSpaceCharacters() {
			common.AbortWithException(c, SpaceCharacterForUsernameException())

			return
		}

		if username.IsTooShort() {
			common.AbortWithException(c, TooShortUsernameException())

			return
		}

		if username.IsTooLong() {
			common.AbortWithException(c, TooLongUsernameException())

			return
		}

		password := entities.NewPassword(body.Password)

		if password.HasSpaceCharacters() {
			common.AbortWithException(c, SpaceCharacterForPasswordException())

			return
		}

		if password.IsTooShort() {
			common.AbortWithException(c, TooShortPasswordException())

			return
		}

		if password.IsTooLong() {
			common.AbortWithException(c, TooLongPasswordException())

			return
		}

		if repositories.UsernameExists(mo.DB, username) {
			common.AbortWithException(c, ConflictUsernameException())

			return
		}

		user := entities.NewUser(username, password)

		repositories.SaveUser(mo.DB, user)

		result := repositories.FindUserByUsername(mo.DB, username)

		c.JSON(http.StatusCreated, &CreateUserResponseBody{
			Id:        result.Id(),
			Username:  result.Username().ToString(),
			CreatedAt: result.CreatedAt(),
		})
	})
}
