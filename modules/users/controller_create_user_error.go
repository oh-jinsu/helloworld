package users

import (
	"net/http"

	"github.com/oh-jinsu/helloworld/modules/common"
)

const (
	CONFLICT_USERNAME = 1001 + iota
	DISALLOW_KOREAN_CHARACTER
	DISALLOW_SPECIAL_CHARACTER
	DISALLOW_SPACE_CHARACTER_FOR_USERNAME
	TOO_SHORT_USERNAME
	TOO_LONG_USERNAME
	TOO_SHORT_PASSWORD
	TOO_LONG_PASSWORD
	DISALLOW_SPACE_CHARACTER_FOR_PASSWORD
)

func ConflictUsernameException() common.Exception {
	return common.NewException(http.StatusConflict, CONFLICT_USERNAME, "이미 존재하는 이름입니다.")
}

func DisallowKoreanCharacter() common.Exception {
	return common.NewException(http.StatusBadRequest, DISALLOW_KOREAN_CHARACTER, "이름에 한글 자음을 입력할 수 없습니다.")
}

func DisallowSpecialCharacter() common.Exception {
	return common.NewException(http.StatusBadRequest, DISALLOW_SPECIAL_CHARACTER, "이름에 특수문자를 입력할 수 없습니다.")
}

func DisallowSpaceCharacterForUsername() common.Exception {
	return common.NewException(http.StatusBadRequest, DISALLOW_SPACE_CHARACTER_FOR_USERNAME, "이름에 공백을 입력할 수 없습니다.")
}

func TooShortUsername() common.Exception {
	return common.NewException(http.StatusBadRequest, TOO_SHORT_USERNAME, "이름이 너무 짧습니다.")
}

func TooLongUsername() common.Exception {
	return common.NewException(http.StatusBadRequest, TOO_LONG_USERNAME, "이름이 너무 깁니다.")
}

func TooShortPassword() common.Exception {
	return common.NewException(http.StatusBadRequest, TOO_SHORT_PASSWORD, "비밀번호가 너무 짧습니다.")
}

func TooLongPassword() common.Exception {
	return common.NewException(http.StatusBadRequest, TOO_LONG_PASSWORD, "비밀번호가 너무 깁니다.")
}

func DisAllowSpaceCharacterForPassword() common.Exception {
	return common.NewException(http.StatusBadRequest, DISALLOW_SPACE_CHARACTER_FOR_PASSWORD, "비밀번호에 공백을 입력할 수 없습니다.")
}
