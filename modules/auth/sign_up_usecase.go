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
			common.AbortWithException(c, common.BadRequestException())

			return
		}

		username := &entities.Username{Value: body.Username}

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

		password := &entities.Password{Value: body.Password}

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

		if models.UsernameExists(mo.Db, username.Value) {
			common.AbortWithException(c, ConflictUsernameException())

			return
		}

		models.SaveUser(mo.Db, username.Value, password.Value)

		user, err := models.FindUserByUsername(mo.Db, username.Value)

		if err != nil {
			common.AbortWithException(c, FailedToFindUserException())

			return
		}

		c.JSON(http.StatusCreated, &SignUpResponseBody{
			Id:        user.Id,
			Username:  user.Username.Value,
			CreatedAt: user.CreatedAt,
		})
	})
}
