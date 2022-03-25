package users

import (
	"net/http"
	"regexp"
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
			common.AbortWithException(c, common.BadRequestException())

			return
		}

		username := body.Username

		if matched, _ := regexp.MatchString("[ㄱ-ㅎ]", username); matched {
			common.AbortWithException(c, DisallowKoreanCharacter())

			return
		}

		if matched, _ := regexp.MatchString("[\\{\\}\\[\\]\\/?.,;:|\\)*~`!^\\-_+<>@\\#$%&\\\\\\=\\(\\'\"]", username); matched {
			common.AbortWithException(c, DisallowSpecialCharacter())

			return
		}

		if matched, _ := regexp.MatchString("\\s", username); matched {
			common.AbortWithException(c, DisallowSpaceCharacterForUsername())

			return
		}

		koreanLength := len(regexp.MustCompile("[가-힣]").FindAllString(username, -1))

		otherLength := len(regexp.MustCompile("[A-Za-z0-9]").FindAllString(username, -1))

		if koreanLength*2+otherLength < 4 {
			common.AbortWithException(c, TooShortUsername())

			return
		}

		if koreanLength*2+otherLength > 16 {
			common.AbortWithException(c, TooLongUsername())

			return
		}

		password := body.Password

		if len(password) < 8 {
			common.AbortWithException(c, TooShortPassword())

			return
		}

		if len(password) > 24 {
			common.AbortWithException(c, TooLongPassword())

			return
		}

		if matched, _ := regexp.MatchString("\\s", password); matched {
			common.AbortWithException(c, DisAllowSpaceCharacterForPassword())

			return
		}

		result := &models.User{}

		if err := mo.DB.Where("username = ?", body.Username).First(result).Error; err == nil {
			common.AbortWithException(c, ConflictUsernameException())

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
